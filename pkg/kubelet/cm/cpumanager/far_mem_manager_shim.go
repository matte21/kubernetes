package cpumanager

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/kubelet/cm/containermap"
	"k8s.io/kubernetes/pkg/kubelet/cm/cpumanager/state"
	"k8s.io/kubernetes/pkg/kubelet/cm/farmemtopologymanager"
	"k8s.io/kubernetes/pkg/kubelet/cm/topologymanager"
	"k8s.io/kubernetes/pkg/kubelet/config"
	"k8s.io/kubernetes/pkg/kubelet/status"
	"k8s.io/utils/cpuset"
)

type farMemManagerShim struct {
	m *farmemtopologymanager.Manager
}

func NewFarMemBasedCPUManager(m *farmemtopologymanager.Manager) Manager {
	return &farMemManagerShim{m: m}
}

var _ Manager = &farMemManagerShim{}

// AddContainer implements Manager.
func (f *farMemManagerShim) AddContainer(p *v1.Pod, c *v1.Container, containerID string) {
	// NOP because it gets called directly on the far memory manager.
}

// Allocate implements Manager.
func (f *farMemManagerShim) Allocate(pod *v1.Pod, container *v1.Container) error {
	panic("unimplemented")
}

// GetAllCPUs implements Manager.
func (f *farMemManagerShim) GetAllCPUs() cpuset.CPUSet {
	return f.m.GetAllCPUs()
}

// GetAllocatableCPUs implements Manager.
func (f *farMemManagerShim) GetAllocatableCPUs() cpuset.CPUSet {
	return f.m.GetAllocatableCPUs()
}

// GetCPUAffinity implements Manager.
func (f *farMemManagerShim) GetCPUAffinity(podUID string, containerName string) cpuset.CPUSet {
	return f.m.GetCPUAffinity(podUID, containerName)
}

// GetExclusiveCPUs implements Manager.
func (f *farMemManagerShim) GetExclusiveCPUs(podUID string, containerName string) cpuset.CPUSet {
	return f.m.GetExclusiveCPUs(podUID, containerName)
}

// GetPodTopologyHints implements Manager.
func (f *farMemManagerShim) GetPodTopologyHints(pod *v1.Pod) map[string][]topologymanager.TopologyHint {
	panic("unimplemented")
}

// GetTopologyHints implements Manager.
func (f *farMemManagerShim) GetTopologyHints(*v1.Pod, *v1.Container) map[string][]topologymanager.TopologyHint {
	panic("unimplemented")
}

// RemoveContainer implements Manager.
func (f *farMemManagerShim) RemoveContainer(containerID string) error {
	// NOP because it gets called directly on the far memory manager.
	return nil
}

// Start implements Manager.
func (f *farMemManagerShim) Start(activePods ActivePodsFunc, sourcesReady config.SourcesReady, podStatusProvider status.PodStatusProvider, containerRuntime runtimeService, initialContainers containermap.ContainerMap) error {
	// NOP, we don't need this but external kubelet code might still call it.
	return nil
}

// State implements Manager.
func (f *farMemManagerShim) State() state.Reader {
	// In theory this isn't ever called. I'd rather do it a nop, but I don't know which state to
	// return. Hence we panic. If I ever see this panic being triggered, I'll think about actually
	// returning something.
	panic("unimplemented")
}
