package main

import (
	"github.com/juju/juju/apiserver/params"
	"github.com/prometheus/client_golang/prometheus"
)

type registry struct {
	*prometheus.Registry
	jujuMachine     *prometheus.GaugeVec
	jujuApplication *prometheus.GaugeVec
	jujuUnit        *prometheus.GaugeVec
	jujuSubordinate *prometheus.GaugeVec
	jujuContainer   *prometheus.GaugeVec
}

func newRegistry(model, modelUUID string) *registry {
	modelLabels := prometheus.Labels{
		"model":      model,
		"model_uuid": modelUUID,
	}
	r := &registry{
		Registry: prometheus.NewRegistry(),
		jujuMachine: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "juju_machine",
				Help: "Juju machine",
			},
			[]string{"model", "model_uuid", "dns_name", "id", "instance_status", "agent_status"},
		).MustCurryWith(modelLabels),
		jujuApplication: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "juju_application",
				Help: "Juju application",
			},
			[]string{"model", "model_uuid", "name", "status"},
		).MustCurryWith(modelLabels),
		jujuUnit: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "juju_unit",
				Help: "Juju unit",
			},
			[]string{"model", "model_uuid", "name", "application_name", "workload_status", "agent_status"},
		).MustCurryWith(modelLabels),
		jujuSubordinate: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "juju_subordinate",
				Help: "Juju subordinate",
			},
			[]string{"model", "model_uuid", "name", "subordinate_to", "application_name", "workload_status", "agent_status"},
		).MustCurryWith(modelLabels),
		jujuContainer: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "juju_container_status",
				Help: "Juju container",
			},
			[]string{"model", "model_uuid", "id", "dns_name", "parent_id", "parent_dns_name", "instance_status", "agent_status"},
		).MustCurryWith(modelLabels),
	}

	r.MustRegister(
		r.jujuApplication,
		r.jujuUnit,
		r.jujuMachine,
		r.jujuSubordinate,
		r.jujuContainer,
	)
	return r
}

func (r *registry) parseStatus(status *params.FullStatus) {
	for applicationName, application := range status.Applications {
		r.jujuApplication.With(prometheus.Labels{
			"name":   applicationName,
			"status": application.Status.Status,
		}).Set(checkStatus(application.Status.Status, []string{"active"}))

		for unitName, unit := range application.Units {
			r.jujuUnit.With(prometheus.Labels{
				"name":             unitName,
				"agent_status":     unit.AgentStatus.Status,
				"workload_status":  unit.WorkloadStatus.Status,
				"application_name": applicationName,
			}).Set(checkStatus(unit.WorkloadStatus.Status, []string{"active", "maintenance"}))

			for subName, sub := range unit.Subordinates {
				r.jujuSubordinate.With(prometheus.Labels{
					"name":             subName,
					"subordinate_to":   unitName,
					"agent_status":     sub.AgentStatus.Status,
					"workload_status":  sub.WorkloadStatus.Status,
					"application_name": applicationName,
				}).Set(checkStatus(sub.WorkloadStatus.Status, []string{"active", "maintenance"}))
			}
		}
	}
	for machineName, machine := range status.Machines {
			"dns_name":        machine.DNSName,
			"id":              machineName,
			"instance_status": machine.InstanceStatus.Status,
			"agent_status":    machine.AgentStatus.Status,
		}).Set(checkStatus(machine.InstanceStatus.Status, []string{"running"}))

		if len(machine.Containers) == 0 {
			continue
		}

		for containerName, container := range machine.Containers {
			r.jujuContainer.With(prometheus.Labels{
				"dns_name":        container.DNSName,
				"id":              containerName,
				"instance_status": container.InstanceStatus.Status,
				"agent_status":    container.AgentStatus.Status,
				"parent_id":       machineName,
				"parent_dns_name": machine.DNSName,
			}).Set(checkStatus(machine.InstanceStatus.Status, []string{"running"}))
		}
	}
}

func checkStatus(status string, accepted []string) float64 {
	if len(accepted) == 0 {
		return 1
	}
	for _, ok := range accepted {
		if status == ok {
			return 1
		}
	}
	return 0
}
