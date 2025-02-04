/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LinkedServicePostgreSQLInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The description for the Data Factory Linked Service PostgreSQL.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServicePostgreSQLObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service PostgreSQL.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory PostgreSQL Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServicePostgreSQLParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with PostgreSQL.
	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service PostgreSQL.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// LinkedServicePostgreSQLSpec defines the desired state of LinkedServicePostgreSQL
type LinkedServicePostgreSQLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServicePostgreSQLParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider LinkedServicePostgreSQLInitParameters `json:"initProvider,omitempty"`
}

// LinkedServicePostgreSQLStatus defines the observed state of LinkedServicePostgreSQL.
type LinkedServicePostgreSQLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServicePostgreSQLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServicePostgreSQL is the Schema for the LinkedServicePostgreSQLs API. Manages a Linked Service (connection) between PostgreSQL and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServicePostgreSQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionString) || has(self.initProvider.connectionString)",message="connectionString is a required parameter"
	Spec   LinkedServicePostgreSQLSpec   `json:"spec"`
	Status LinkedServicePostgreSQLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServicePostgreSQLList contains a list of LinkedServicePostgreSQLs
type LinkedServicePostgreSQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServicePostgreSQL `json:"items"`
}

// Repository type metadata.
var (
	LinkedServicePostgreSQL_Kind             = "LinkedServicePostgreSQL"
	LinkedServicePostgreSQL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServicePostgreSQL_Kind}.String()
	LinkedServicePostgreSQL_KindAPIVersion   = LinkedServicePostgreSQL_Kind + "." + CRDGroupVersion.String()
	LinkedServicePostgreSQL_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServicePostgreSQL_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServicePostgreSQL{}, &LinkedServicePostgreSQLList{})
}
