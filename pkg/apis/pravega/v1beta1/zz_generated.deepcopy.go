// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthHandlerSpec) DeepCopyInto(out *AuthHandlerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthHandlerSpec.
func (in *AuthHandlerSpec) DeepCopy() *AuthHandlerSpec {
	if in == nil {
		return nil
	}
	out := new(AuthHandlerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthImplementationSpec) DeepCopyInto(out *AuthImplementationSpec) {
	*out = *in
	if in.AuthHandlers != nil {
		in, out := &in.AuthHandlers, &out.AuthHandlers
		*out = make([]AuthHandlerSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthImplementationSpec.
func (in *AuthImplementationSpec) DeepCopy() *AuthImplementationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthImplementationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationParameters) DeepCopyInto(out *AuthenticationParameters) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationParameters.
func (in *AuthenticationParameters) DeepCopy() *AuthenticationParameters {
	if in == nil {
		return nil
	}
	out := new(AuthenticationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.ExternalAccess != nil {
		in, out := &in.ExternalAccess, &out.ExternalAccess
		*out = new(ExternalAccess)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(AuthenticationParameters)
		**out = **in
	}
	if in.Pravega != nil {
		in, out := &in.Pravega, &out.Pravega
		*out = new(PravegaSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		copy(*out, *in)
	}
	if in.VersionHistory != nil {
		in, out := &in.VersionHistory, &out.VersionHistory
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Members.DeepCopyInto(&out.Members)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSpec) DeepCopyInto(out *CustomSpec) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSpec.
func (in *CustomSpec) DeepCopy() *CustomSpec {
	if in == nil {
		return nil
	}
	out := new(CustomSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECSSpec) DeepCopyInto(out *ECSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECSSpec.
func (in *ECSSpec) DeepCopy() *ECSSpec {
	if in == nil {
		return nil
	}
	out := new(ECSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAccess) DeepCopyInto(out *ExternalAccess) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAccess.
func (in *ExternalAccess) DeepCopy() *ExternalAccess {
	if in == nil {
		return nil
	}
	out := new(ExternalAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemSpec) DeepCopyInto(out *FileSystemSpec) {
	*out = *in
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(v1.PersistentVolumeClaimVolumeSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemSpec.
func (in *FileSystemSpec) DeepCopy() *FileSystemSpec {
	if in == nil {
		return nil
	}
	out := new(FileSystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HDFSSpec) DeepCopyInto(out *HDFSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HDFSSpec.
func (in *HDFSSpec) DeepCopy() *HDFSSpec {
	if in == nil {
		return nil
	}
	out := new(HDFSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfluxDBSecret) DeepCopyInto(out *InfluxDBSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfluxDBSecret.
func (in *InfluxDBSecret) DeepCopy() *InfluxDBSecret {
	if in == nil {
		return nil
	}
	out := new(InfluxDBSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LongTermStorageSpec) DeepCopyInto(out *LongTermStorageSpec) {
	*out = *in
	if in.FileSystem != nil {
		in, out := &in.FileSystem, &out.FileSystem
		*out = new(FileSystemSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Ecs != nil {
		in, out := &in.Ecs, &out.Ecs
		*out = new(ECSSpec)
		**out = **in
	}
	if in.Hdfs != nil {
		in, out := &in.Hdfs, &out.Hdfs
		*out = new(HDFSSpec)
		**out = **in
	}
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LongTermStorageSpec.
func (in *LongTermStorageSpec) DeepCopy() *LongTermStorageSpec {
	if in == nil {
		return nil
	}
	out := new(LongTermStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembersStatus) DeepCopyInto(out *MembersStatus) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Unready != nil {
		in, out := &in.Unready, &out.Unready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembersStatus.
func (in *MembersStatus) DeepCopy() *MembersStatus {
	if in == nil {
		return nil
	}
	out := new(MembersStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PravegaCluster) DeepCopyInto(out *PravegaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PravegaCluster.
func (in *PravegaCluster) DeepCopy() *PravegaCluster {
	if in == nil {
		return nil
	}
	out := new(PravegaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PravegaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PravegaClusterList) DeepCopyInto(out *PravegaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PravegaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PravegaClusterList.
func (in *PravegaClusterList) DeepCopy() *PravegaClusterList {
	if in == nil {
		return nil
	}
	out := new(PravegaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PravegaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PravegaSpec) DeepCopyInto(out *PravegaSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ControllerJvmOptions != nil {
		in, out := &in.ControllerJvmOptions, &out.ControllerJvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SegmentStoreJVMOptions != nil {
		in, out := &in.SegmentStoreJVMOptions, &out.SegmentStoreJVMOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CacheVolumeClaimTemplate != nil {
		in, out := &in.CacheVolumeClaimTemplate, &out.CacheVolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LongTermStorage != nil {
		in, out := &in.LongTermStorage, &out.LongTermStorage
		*out = new(LongTermStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerResources != nil {
		in, out := &in.ControllerResources, &out.ControllerResources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.SegmentStoreResources != nil {
		in, out := &in.SegmentStoreResources, &out.SegmentStoreResources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.SegmentStoreSecret != nil {
		in, out := &in.SegmentStoreSecret, &out.SegmentStoreSecret
		*out = new(SegmentStoreSecret)
		**out = **in
	}
	if in.ControllerServiceAnnotations != nil {
		in, out := &in.ControllerServiceAnnotations, &out.ControllerServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ControllerPodLabels != nil {
		in, out := &in.ControllerPodLabels, &out.ControllerPodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ControllerPodAnnotations != nil {
		in, out := &in.ControllerPodAnnotations, &out.ControllerPodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SegmentStoreServiceAnnotations != nil {
		in, out := &in.SegmentStoreServiceAnnotations, &out.SegmentStoreServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SegmentStorePodLabels != nil {
		in, out := &in.SegmentStorePodLabels, &out.SegmentStorePodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SegmentStorePodAnnotations != nil {
		in, out := &in.SegmentStorePodAnnotations, &out.SegmentStorePodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SegmentStoreSecurityContext != nil {
		in, out := &in.SegmentStoreSecurityContext, &out.SegmentStoreSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerSecurityContext != nil {
		in, out := &in.ControllerSecurityContext, &out.ControllerSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerPodAffinity != nil {
		in, out := &in.ControllerPodAffinity, &out.ControllerPodAffinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.SegmentStorePodAffinity != nil {
		in, out := &in.SegmentStorePodAffinity, &out.SegmentStorePodAffinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerInitContainers != nil {
		in, out := &in.ControllerInitContainers, &out.ControllerInitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SegmentStoreInitContainers != nil {
		in, out := &in.SegmentStoreInitContainers, &out.SegmentStoreInitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthImplementations != nil {
		in, out := &in.AuthImplementations, &out.AuthImplementations
		*out = new(AuthImplementationSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.InfluxDBSecret != nil {
		in, out := &in.InfluxDBSecret, &out.InfluxDBSecret
		*out = new(InfluxDBSecret)
		**out = **in
	}
	if in.ControllerProbes != nil {
		in, out := &in.ControllerProbes, &out.ControllerProbes
		*out = new(Probes)
		(*in).DeepCopyInto(*out)
	}
	if in.SegmentStoreProbes != nil {
		in, out := &in.SegmentStoreProbes, &out.SegmentStoreProbes
		*out = new(Probes)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PravegaSpec.
func (in *PravegaSpec) DeepCopy() *PravegaSpec {
	if in == nil {
		return nil
	}
	out := new(PravegaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probe) DeepCopyInto(out *Probe) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probe.
func (in *Probe) DeepCopy() *Probe {
	if in == nil {
		return nil
	}
	out := new(Probe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probes) DeepCopyInto(out *Probes) {
	*out = *in
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(Probe)
		**out = **in
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(Probe)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probes.
func (in *Probes) DeepCopy() *Probes {
	if in == nil {
		return nil
	}
	out := new(Probes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SegmentStoreSecret) DeepCopyInto(out *SegmentStoreSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SegmentStoreSecret.
func (in *SegmentStoreSecret) DeepCopy() *SegmentStoreSecret {
	if in == nil {
		return nil
	}
	out := new(SegmentStoreSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticTLS) DeepCopyInto(out *StaticTLS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticTLS.
func (in *StaticTLS) DeepCopy() *StaticTLS {
	if in == nil {
		return nil
	}
	out := new(StaticTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSPolicy) DeepCopyInto(out *TLSPolicy) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(StaticTLS)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSPolicy.
func (in *TLSPolicy) DeepCopy() *TLSPolicy {
	if in == nil {
		return nil
	}
	out := new(TLSPolicy)
	in.DeepCopyInto(out)
	return out
}
