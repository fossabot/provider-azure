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

type PolicyInitParameters struct {
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	EnvironmentType *string `json:"environmentType,omitempty" tf:"environment_type,omitempty"`
}

type PolicyObservation struct {
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	EnvironmentType *string `json:"environmentType,omitempty" tf:"environment_type,omitempty"`
}

type PolicyParameters struct {

	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// +kubebuilder:validation:Optional
	EnvironmentType *string `json:"environmentType,omitempty" tf:"environment_type,omitempty"`
}

type ProviderInitParameters struct {

	// The Azure Region where the Attestation Provider should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	OpenEnclavePolicyBase64 *string `json:"openEnclavePolicyBase64,omitempty" tf:"open_enclave_policy_base64,omitempty"`

	Policy []PolicyInitParameters `json:"policy,omitempty" tf:"policy,omitempty"`

	// A valid X.509 certificate (Section 4 of RFC4648). Changing this forces a new resource to be created.
	PolicySigningCertificateData *string `json:"policySigningCertificateData,omitempty" tf:"policy_signing_certificate_data,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	SevSnpPolicyBase64 *string `json:"sevSnpPolicyBase64,omitempty" tf:"sev_snp_policy_base64,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	SgxEnclavePolicyBase64 *string `json:"sgxEnclavePolicyBase64,omitempty" tf:"sgx_enclave_policy_base64,omitempty"`

	// A mapping of tags which should be assigned to the Attestation Provider.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	TpmPolicyBase64 *string `json:"tpmPolicyBase64,omitempty" tf:"tpm_policy_base64,omitempty"`
}

type ProviderObservation struct {

	// The URI of the Attestation Service.
	AttestationURI *string `json:"attestationUri,omitempty" tf:"attestation_uri,omitempty"`

	// The ID of the Attestation Provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Attestation Provider should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	OpenEnclavePolicyBase64 *string `json:"openEnclavePolicyBase64,omitempty" tf:"open_enclave_policy_base64,omitempty"`

	Policy []PolicyObservation `json:"policy,omitempty" tf:"policy,omitempty"`

	// A valid X.509 certificate (Section 4 of RFC4648). Changing this forces a new resource to be created.
	PolicySigningCertificateData *string `json:"policySigningCertificateData,omitempty" tf:"policy_signing_certificate_data,omitempty"`

	// The name of the Resource Group where the attestation provider should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	SevSnpPolicyBase64 *string `json:"sevSnpPolicyBase64,omitempty" tf:"sev_snp_policy_base64,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	SgxEnclavePolicyBase64 *string `json:"sgxEnclavePolicyBase64,omitempty" tf:"sgx_enclave_policy_base64,omitempty"`

	// A mapping of tags which should be assigned to the Attestation Provider.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	TpmPolicyBase64 *string `json:"tpmPolicyBase64,omitempty" tf:"tpm_policy_base64,omitempty"`

	// Trust model used for the Attestation Service.
	TrustModel *string `json:"trustModel,omitempty" tf:"trust_model,omitempty"`
}

type ProviderParameters struct {

	// The Azure Region where the Attestation Provider should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	// +kubebuilder:validation:Optional
	OpenEnclavePolicyBase64 *string `json:"openEnclavePolicyBase64,omitempty" tf:"open_enclave_policy_base64,omitempty"`

	// +kubebuilder:validation:Optional
	Policy []PolicyParameters `json:"policy,omitempty" tf:"policy,omitempty"`

	// A valid X.509 certificate (Section 4 of RFC4648). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PolicySigningCertificateData *string `json:"policySigningCertificateData,omitempty" tf:"policy_signing_certificate_data,omitempty"`

	// The name of the Resource Group where the attestation provider should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	// +kubebuilder:validation:Optional
	SevSnpPolicyBase64 *string `json:"sevSnpPolicyBase64,omitempty" tf:"sev_snp_policy_base64,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	// +kubebuilder:validation:Optional
	SgxEnclavePolicyBase64 *string `json:"sgxEnclavePolicyBase64,omitempty" tf:"sgx_enclave_policy_base64,omitempty"`

	// A mapping of tags which should be assigned to the Attestation Provider.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the base64 URI Encoded RFC 7519 JWT that should be used for the Attestation Policy.
	// +kubebuilder:validation:Optional
	TpmPolicyBase64 *string `json:"tpmPolicyBase64,omitempty" tf:"tpm_policy_base64,omitempty"`
}

// ProviderSpec defines the desired state of Provider
type ProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProviderParameters `json:"forProvider"`
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
	InitProvider ProviderInitParameters `json:"initProvider,omitempty"`
}

// ProviderStatus defines the observed state of Provider.
type ProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Provider is the Schema for the Providers API. Manages an Attestation Provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Provider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   ProviderSpec   `json:"spec"`
	Status ProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderList contains a list of Providers
type ProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Provider `json:"items"`
}

// Repository type metadata.
var (
	Provider_Kind             = "Provider"
	Provider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Provider_Kind}.String()
	Provider_KindAPIVersion   = Provider_Kind + "." + CRDGroupVersion.String()
	Provider_GroupVersionKind = CRDGroupVersion.WithKind(Provider_Kind)
)

func init() {
	SchemeBuilder.Register(&Provider{}, &ProviderList{})
}
