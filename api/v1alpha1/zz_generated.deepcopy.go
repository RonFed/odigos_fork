//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/keyval-dev/odigos/common"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collector) DeepCopyInto(out *Collector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collector.
func (in *Collector) DeepCopy() *Collector {
	if in == nil {
		return nil
	}
	out := new(Collector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Collector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorList) DeepCopyInto(out *CollectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Collector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorList.
func (in *CollectorList) DeepCopy() *CollectorList {
	if in == nil {
		return nil
	}
	out := new(CollectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorSpec) DeepCopyInto(out *CollectorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorSpec.
func (in *CollectorSpec) DeepCopy() *CollectorSpec {
	if in == nil {
		return nil
	}
	out := new(CollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorStatus) DeepCopyInto(out *CollectorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorStatus.
func (in *CollectorStatus) DeepCopy() *CollectorStatus {
	if in == nil {
		return nil
	}
	out := new(CollectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogData) DeepCopyInto(out *DatadogData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogData.
func (in *DatadogData) DeepCopy() *DatadogData {
	if in == nil {
		return nil
	}
	out := new(DatadogData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Destination) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationData) DeepCopyInto(out *DestinationData) {
	*out = *in
	if in.Grafana != nil {
		in, out := &in.Grafana, &out.Grafana
		*out = new(GrafanaData)
		**out = **in
	}
	if in.Honeycomb != nil {
		in, out := &in.Honeycomb, &out.Honeycomb
		*out = new(HoneycombData)
		**out = **in
	}
	if in.Datadog != nil {
		in, out := &in.Datadog, &out.Datadog
		*out = new(DatadogData)
		**out = **in
	}
	if in.NewRelic != nil {
		in, out := &in.NewRelic, &out.NewRelic
		*out = new(NewRelicData)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationData.
func (in *DestinationData) DeepCopy() *DestinationData {
	if in == nil {
		return nil
	}
	out := new(DestinationData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationList) DeepCopyInto(out *DestinationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Destination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationList.
func (in *DestinationList) DeepCopy() *DestinationList {
	if in == nil {
		return nil
	}
	out := new(DestinationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DestinationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationSpec) DeepCopyInto(out *DestinationSpec) {
	*out = *in
	in.Data.DeepCopyInto(&out.Data)
	out.SecretRef = in.SecretRef
	if in.Signals != nil {
		in, out := &in.Signals, &out.Signals
		*out = make([]common.ObservabilitySignal, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationSpec.
func (in *DestinationSpec) DeepCopy() *DestinationSpec {
	if in == nil {
		return nil
	}
	out := new(DestinationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationStatus) DeepCopyInto(out *DestinationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationStatus.
func (in *DestinationStatus) DeepCopy() *DestinationStatus {
	if in == nil {
		return nil
	}
	out := new(DestinationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaData) DeepCopyInto(out *GrafanaData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaData.
func (in *GrafanaData) DeepCopy() *GrafanaData {
	if in == nil {
		return nil
	}
	out := new(GrafanaData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HoneycombData) DeepCopyInto(out *HoneycombData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HoneycombData.
func (in *HoneycombData) DeepCopy() *HoneycombData {
	if in == nil {
		return nil
	}
	out := new(HoneycombData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstrumentedApplication) DeepCopyInto(out *InstrumentedApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstrumentedApplication.
func (in *InstrumentedApplication) DeepCopy() *InstrumentedApplication {
	if in == nil {
		return nil
	}
	out := new(InstrumentedApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstrumentedApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstrumentedApplicationList) DeepCopyInto(out *InstrumentedApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstrumentedApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstrumentedApplicationList.
func (in *InstrumentedApplicationList) DeepCopy() *InstrumentedApplicationList {
	if in == nil {
		return nil
	}
	out := new(InstrumentedApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstrumentedApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstrumentedApplicationSpec) DeepCopyInto(out *InstrumentedApplicationSpec) {
	*out = *in
	if in.Languages != nil {
		in, out := &in.Languages, &out.Languages
		*out = make([]common.LanguageByContainer, len(*in))
		copy(*out, *in)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstrumentedApplicationSpec.
func (in *InstrumentedApplicationSpec) DeepCopy() *InstrumentedApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(InstrumentedApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstrumentedApplicationStatus) DeepCopyInto(out *InstrumentedApplicationStatus) {
	*out = *in
	out.LangDetection = in.LangDetection
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstrumentedApplicationStatus.
func (in *InstrumentedApplicationStatus) DeepCopy() *InstrumentedApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(InstrumentedApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LangDetectionStatus) DeepCopyInto(out *LangDetectionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LangDetectionStatus.
func (in *LangDetectionStatus) DeepCopy() *LangDetectionStatus {
	if in == nil {
		return nil
	}
	out := new(LangDetectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NewRelicData) DeepCopyInto(out *NewRelicData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NewRelicData.
func (in *NewRelicData) DeepCopy() *NewRelicData {
	if in == nil {
		return nil
	}
	out := new(NewRelicData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OdigosConfiguration) DeepCopyInto(out *OdigosConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OdigosConfiguration.
func (in *OdigosConfiguration) DeepCopy() *OdigosConfiguration {
	if in == nil {
		return nil
	}
	out := new(OdigosConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OdigosConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OdigosConfigurationList) DeepCopyInto(out *OdigosConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OdigosConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OdigosConfigurationList.
func (in *OdigosConfigurationList) DeepCopy() *OdigosConfigurationList {
	if in == nil {
		return nil
	}
	out := new(OdigosConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OdigosConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OdigosConfigurationSpec) DeepCopyInto(out *OdigosConfigurationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OdigosConfigurationSpec.
func (in *OdigosConfigurationSpec) DeepCopy() *OdigosConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(OdigosConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}