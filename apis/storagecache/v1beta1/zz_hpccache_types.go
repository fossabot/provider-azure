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

type AccessRuleInitParameters struct {

	// The access level for this rule. Possible values are: rw, ro, no.
	Access *string `json:"access,omitempty" tf:"access,omitempty"`

	// The anonymous GID used when root_squash_enabled is true.
	AnonymousGID *float64 `json:"anonymousGid,omitempty" tf:"anonymous_gid,omitempty"`

	// The anonymous UID used when root_squash_enabled is true.
	AnonymousUID *float64 `json:"anonymousUid,omitempty" tf:"anonymous_uid,omitempty"`

	// The filter applied to the scope for this rule. The filter's format depends on its scope: default scope matches all clients and has no filter value; network scope takes a CIDR format; host takes an IP address or fully qualified domain name. If a client does not match any filter rule and there is no default rule, access is denied.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Whether to enable root squash?
	RootSquashEnabled *bool `json:"rootSquashEnabled,omitempty" tf:"root_squash_enabled,omitempty"`

	// The scope of this rule. The scope and (potentially) the filter determine which clients match the rule. Possible values are: default, network, host.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Whether allow access to subdirectories under the root export?
	SubmountAccessEnabled *bool `json:"submountAccessEnabled,omitempty" tf:"submount_access_enabled,omitempty"`

	// Whether SUID is allowed?
	SuidEnabled *bool `json:"suidEnabled,omitempty" tf:"suid_enabled,omitempty"`
}

type AccessRuleObservation struct {

	// The access level for this rule. Possible values are: rw, ro, no.
	Access *string `json:"access,omitempty" tf:"access,omitempty"`

	// The anonymous GID used when root_squash_enabled is true.
	AnonymousGID *float64 `json:"anonymousGid,omitempty" tf:"anonymous_gid,omitempty"`

	// The anonymous UID used when root_squash_enabled is true.
	AnonymousUID *float64 `json:"anonymousUid,omitempty" tf:"anonymous_uid,omitempty"`

	// The filter applied to the scope for this rule. The filter's format depends on its scope: default scope matches all clients and has no filter value; network scope takes a CIDR format; host takes an IP address or fully qualified domain name. If a client does not match any filter rule and there is no default rule, access is denied.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Whether to enable root squash?
	RootSquashEnabled *bool `json:"rootSquashEnabled,omitempty" tf:"root_squash_enabled,omitempty"`

	// The scope of this rule. The scope and (potentially) the filter determine which clients match the rule. Possible values are: default, network, host.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Whether allow access to subdirectories under the root export?
	SubmountAccessEnabled *bool `json:"submountAccessEnabled,omitempty" tf:"submount_access_enabled,omitempty"`

	// Whether SUID is allowed?
	SuidEnabled *bool `json:"suidEnabled,omitempty" tf:"suid_enabled,omitempty"`
}

type AccessRuleParameters struct {

	// The access level for this rule. Possible values are: rw, ro, no.
	// +kubebuilder:validation:Optional
	Access *string `json:"access" tf:"access,omitempty"`

	// The anonymous GID used when root_squash_enabled is true.
	// +kubebuilder:validation:Optional
	AnonymousGID *float64 `json:"anonymousGid,omitempty" tf:"anonymous_gid,omitempty"`

	// The anonymous UID used when root_squash_enabled is true.
	// +kubebuilder:validation:Optional
	AnonymousUID *float64 `json:"anonymousUid,omitempty" tf:"anonymous_uid,omitempty"`

	// The filter applied to the scope for this rule. The filter's format depends on its scope: default scope matches all clients and has no filter value; network scope takes a CIDR format; host takes an IP address or fully qualified domain name. If a client does not match any filter rule and there is no default rule, access is denied.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Whether to enable root squash?
	// +kubebuilder:validation:Optional
	RootSquashEnabled *bool `json:"rootSquashEnabled,omitempty" tf:"root_squash_enabled,omitempty"`

	// The scope of this rule. The scope and (potentially) the filter determine which clients match the rule. Possible values are: default, network, host.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope" tf:"scope,omitempty"`

	// Whether allow access to subdirectories under the root export?
	// +kubebuilder:validation:Optional
	SubmountAccessEnabled *bool `json:"submountAccessEnabled,omitempty" tf:"submount_access_enabled,omitempty"`

	// Whether SUID is allowed?
	// +kubebuilder:validation:Optional
	SuidEnabled *bool `json:"suidEnabled,omitempty" tf:"suid_enabled,omitempty"`
}

type BindInitParameters struct {

	// The Bind Distinguished Name (DN) identity to be used in the secure LDAP connection.
	Dn *string `json:"dn,omitempty" tf:"dn,omitempty"`
}

type BindObservation struct {

	// The Bind Distinguished Name (DN) identity to be used in the secure LDAP connection.
	Dn *string `json:"dn,omitempty" tf:"dn,omitempty"`
}

type BindParameters struct {

	// The Bind Distinguished Name (DN) identity to be used in the secure LDAP connection.
	// +kubebuilder:validation:Optional
	Dn *string `json:"dn" tf:"dn,omitempty"`

	// The password of the Active Directory domain administrator.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`
}

type DNSInitParameters struct {

	// The DNS search domain for the HPC Cache.
	SearchDomain *string `json:"searchDomain,omitempty" tf:"search_domain,omitempty"`

	// A list of DNS servers for the HPC Cache. At most three IP(s) are allowed to set.
	Servers []*string `json:"servers,omitempty" tf:"servers,omitempty"`
}

type DNSObservation struct {

	// The DNS search domain for the HPC Cache.
	SearchDomain *string `json:"searchDomain,omitempty" tf:"search_domain,omitempty"`

	// A list of DNS servers for the HPC Cache. At most three IP(s) are allowed to set.
	Servers []*string `json:"servers,omitempty" tf:"servers,omitempty"`
}

type DNSParameters struct {

	// The DNS search domain for the HPC Cache.
	// +kubebuilder:validation:Optional
	SearchDomain *string `json:"searchDomain,omitempty" tf:"search_domain,omitempty"`

	// A list of DNS servers for the HPC Cache. At most three IP(s) are allowed to set.
	// +kubebuilder:validation:Optional
	Servers []*string `json:"servers" tf:"servers,omitempty"`
}

type DefaultAccessPolicyInitParameters struct {

	// One to three access_rule blocks as defined above.
	AccessRule []AccessRuleInitParameters `json:"accessRule,omitempty" tf:"access_rule,omitempty"`
}

type DefaultAccessPolicyObservation struct {

	// One to three access_rule blocks as defined above.
	AccessRule []AccessRuleObservation `json:"accessRule,omitempty" tf:"access_rule,omitempty"`
}

type DefaultAccessPolicyParameters struct {

	// One to three access_rule blocks as defined above.
	// +kubebuilder:validation:Optional
	AccessRule []AccessRuleParameters `json:"accessRule" tf:"access_rule,omitempty"`
}

type DirectoryActiveDirectoryInitParameters struct {

	// The NetBIOS name to assign to the HPC Cache when it joins the Active Directory domain as a server.
	CacheNetbiosName *string `json:"cacheNetbiosName,omitempty" tf:"cache_netbios_name,omitempty"`

	// The primary DNS IP address used to resolve the Active Directory domain controller's FQDN.
	DNSPrimaryIP *string `json:"dnsPrimaryIp,omitempty" tf:"dns_primary_ip,omitempty"`

	// The secondary DNS IP address used to resolve the Active Directory domain controller's FQDN.
	DNSSecondaryIP *string `json:"dnsSecondaryIp,omitempty" tf:"dns_secondary_ip,omitempty"`

	// The fully qualified domain name of the Active Directory domain controller.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The Active Directory domain's NetBIOS name.
	DomainNetbiosName *string `json:"domainNetbiosName,omitempty" tf:"domain_netbios_name,omitempty"`

	// The username of the Active Directory domain administrator.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type DirectoryActiveDirectoryObservation struct {

	// The NetBIOS name to assign to the HPC Cache when it joins the Active Directory domain as a server.
	CacheNetbiosName *string `json:"cacheNetbiosName,omitempty" tf:"cache_netbios_name,omitempty"`

	// The primary DNS IP address used to resolve the Active Directory domain controller's FQDN.
	DNSPrimaryIP *string `json:"dnsPrimaryIp,omitempty" tf:"dns_primary_ip,omitempty"`

	// The secondary DNS IP address used to resolve the Active Directory domain controller's FQDN.
	DNSSecondaryIP *string `json:"dnsSecondaryIp,omitempty" tf:"dns_secondary_ip,omitempty"`

	// The fully qualified domain name of the Active Directory domain controller.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The Active Directory domain's NetBIOS name.
	DomainNetbiosName *string `json:"domainNetbiosName,omitempty" tf:"domain_netbios_name,omitempty"`

	// The username of the Active Directory domain administrator.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type DirectoryActiveDirectoryParameters struct {

	// The NetBIOS name to assign to the HPC Cache when it joins the Active Directory domain as a server.
	// +kubebuilder:validation:Optional
	CacheNetbiosName *string `json:"cacheNetbiosName" tf:"cache_netbios_name,omitempty"`

	// The primary DNS IP address used to resolve the Active Directory domain controller's FQDN.
	// +kubebuilder:validation:Optional
	DNSPrimaryIP *string `json:"dnsPrimaryIp" tf:"dns_primary_ip,omitempty"`

	// The secondary DNS IP address used to resolve the Active Directory domain controller's FQDN.
	// +kubebuilder:validation:Optional
	DNSSecondaryIP *string `json:"dnsSecondaryIp,omitempty" tf:"dns_secondary_ip,omitempty"`

	// The fully qualified domain name of the Active Directory domain controller.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// The Active Directory domain's NetBIOS name.
	// +kubebuilder:validation:Optional
	DomainNetbiosName *string `json:"domainNetbiosName" tf:"domain_netbios_name,omitempty"`

	// The password of the Active Directory domain administrator.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username of the Active Directory domain administrator.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type DirectoryFlatFileInitParameters struct {

	// The URI of the file containing group information (/etc/group file format in Unix-like OS).
	GroupFileURI *string `json:"groupFileUri,omitempty" tf:"group_file_uri,omitempty"`

	// The URI of the file containing user information (/etc/passwd file format in Unix-like OS).
	PasswordFileURI *string `json:"passwordFileUri,omitempty" tf:"password_file_uri,omitempty"`
}

type DirectoryFlatFileObservation struct {

	// The URI of the file containing group information (/etc/group file format in Unix-like OS).
	GroupFileURI *string `json:"groupFileUri,omitempty" tf:"group_file_uri,omitempty"`

	// The URI of the file containing user information (/etc/passwd file format in Unix-like OS).
	PasswordFileURI *string `json:"passwordFileUri,omitempty" tf:"password_file_uri,omitempty"`
}

type DirectoryFlatFileParameters struct {

	// The URI of the file containing group information (/etc/group file format in Unix-like OS).
	// +kubebuilder:validation:Optional
	GroupFileURI *string `json:"groupFileUri" tf:"group_file_uri,omitempty"`

	// The URI of the file containing user information (/etc/passwd file format in Unix-like OS).
	// +kubebuilder:validation:Optional
	PasswordFileURI *string `json:"passwordFileUri" tf:"password_file_uri,omitempty"`
}

type DirectoryLdapInitParameters struct {

	// The base distinguished name (DN) for the LDAP domain.
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	// A bind block as defined above.
	Bind []BindInitParameters `json:"bind,omitempty" tf:"bind,omitempty"`

	// The URI of the CA certificate to validate the LDAP secure connection.
	CertificateValidationURI *string `json:"certificateValidationUri,omitempty" tf:"certificate_validation_uri,omitempty"`

	// Whether the certificate should be automatically downloaded. This can be set to true only when certificate_validation_uri is provided.
	DownloadCertificateAutomatically *bool `json:"downloadCertificateAutomatically,omitempty" tf:"download_certificate_automatically,omitempty"`

	// Whether the LDAP connection should be encrypted?
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The FQDN or IP address of the LDAP server.
	Server *string `json:"server,omitempty" tf:"server,omitempty"`
}

type DirectoryLdapObservation struct {

	// The base distinguished name (DN) for the LDAP domain.
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	// A bind block as defined above.
	Bind []BindObservation `json:"bind,omitempty" tf:"bind,omitempty"`

	// The URI of the CA certificate to validate the LDAP secure connection.
	CertificateValidationURI *string `json:"certificateValidationUri,omitempty" tf:"certificate_validation_uri,omitempty"`

	// Whether the certificate should be automatically downloaded. This can be set to true only when certificate_validation_uri is provided.
	DownloadCertificateAutomatically *bool `json:"downloadCertificateAutomatically,omitempty" tf:"download_certificate_automatically,omitempty"`

	// Whether the LDAP connection should be encrypted?
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The FQDN or IP address of the LDAP server.
	Server *string `json:"server,omitempty" tf:"server,omitempty"`
}

type DirectoryLdapParameters struct {

	// The base distinguished name (DN) for the LDAP domain.
	// +kubebuilder:validation:Optional
	BaseDn *string `json:"baseDn" tf:"base_dn,omitempty"`

	// A bind block as defined above.
	// +kubebuilder:validation:Optional
	Bind []BindParameters `json:"bind,omitempty" tf:"bind,omitempty"`

	// The URI of the CA certificate to validate the LDAP secure connection.
	// +kubebuilder:validation:Optional
	CertificateValidationURI *string `json:"certificateValidationUri,omitempty" tf:"certificate_validation_uri,omitempty"`

	// Whether the certificate should be automatically downloaded. This can be set to true only when certificate_validation_uri is provided.
	// +kubebuilder:validation:Optional
	DownloadCertificateAutomatically *bool `json:"downloadCertificateAutomatically,omitempty" tf:"download_certificate_automatically,omitempty"`

	// Whether the LDAP connection should be encrypted?
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The FQDN or IP address of the LDAP server.
	// +kubebuilder:validation:Optional
	Server *string `json:"server" tf:"server,omitempty"`
}

type HPCCacheInitParameters struct {

	// Specifies whether the HPC Cache automatically rotates Encryption Key to the latest version.
	AutomaticallyRotateKeyToLatestEnabled *bool `json:"automaticallyRotateKeyToLatestEnabled,omitempty" tf:"automatically_rotate_key_to_latest_enabled,omitempty"`

	// The size of the HPC Cache, in GB. Possible values are 3072, 6144, 12288, 21623, 24576, 43246, 49152 and 86491. Changing this forces a new resource to be created.
	CacheSizeInGb *float64 `json:"cacheSizeInGb,omitempty" tf:"cache_size_in_gb,omitempty"`

	// A dns block as defined below.
	DNS []DNSInitParameters `json:"dns,omitempty" tf:"dns,omitempty"`

	// A default_access_policy block as defined below.
	DefaultAccessPolicy []DefaultAccessPolicyInitParameters `json:"defaultAccessPolicy,omitempty" tf:"default_access_policy,omitempty"`

	// A directory_active_directory block as defined below.
	DirectoryActiveDirectory []DirectoryActiveDirectoryInitParameters `json:"directoryActiveDirectory,omitempty" tf:"directory_active_directory,omitempty"`

	// A directory_flat_file block as defined below.
	DirectoryFlatFile []DirectoryFlatFileInitParameters `json:"directoryFlatFile,omitempty" tf:"directory_flat_file,omitempty"`

	// A directory_ldap block as defined below.
	DirectoryLdap []DirectoryLdapInitParameters `json:"directoryLdap,omitempty" tf:"directory_ldap,omitempty"`

	// An identity block as defined below. Changing this forces a new resource to be created.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The ID of the Key Vault Key which should be used to encrypt the data in this HPC Cache.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Specifies the supported Azure Region where the HPC Cache should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The IPv4 maximum transmission unit configured for the subnet of the HPC Cache. Possible values range from 576 - 1500. Defaults to 1500.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The NTP server IP Address or FQDN for the HPC Cache. Defaults to time.windows.com.
	NtpServer *string `json:"ntpServer,omitempty" tf:"ntp_server,omitempty"`

	// The SKU of HPC Cache to use. Possible values are (ReadWrite) - Standard_2G, Standard_4G Standard_8G or (ReadOnly) - Standard_L4_5G, Standard_L9G, and Standard_L16G. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the HPC Cache.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HPCCacheObservation struct {

	// Specifies whether the HPC Cache automatically rotates Encryption Key to the latest version.
	AutomaticallyRotateKeyToLatestEnabled *bool `json:"automaticallyRotateKeyToLatestEnabled,omitempty" tf:"automatically_rotate_key_to_latest_enabled,omitempty"`

	// The size of the HPC Cache, in GB. Possible values are 3072, 6144, 12288, 21623, 24576, 43246, 49152 and 86491. Changing this forces a new resource to be created.
	CacheSizeInGb *float64 `json:"cacheSizeInGb,omitempty" tf:"cache_size_in_gb,omitempty"`

	// A dns block as defined below.
	DNS []DNSObservation `json:"dns,omitempty" tf:"dns,omitempty"`

	// A default_access_policy block as defined below.
	DefaultAccessPolicy []DefaultAccessPolicyObservation `json:"defaultAccessPolicy,omitempty" tf:"default_access_policy,omitempty"`

	// A directory_active_directory block as defined below.
	DirectoryActiveDirectory []DirectoryActiveDirectoryObservation `json:"directoryActiveDirectory,omitempty" tf:"directory_active_directory,omitempty"`

	// A directory_flat_file block as defined below.
	DirectoryFlatFile []DirectoryFlatFileObservation `json:"directoryFlatFile,omitempty" tf:"directory_flat_file,omitempty"`

	// A directory_ldap block as defined below.
	DirectoryLdap []DirectoryLdapObservation `json:"directoryLdap,omitempty" tf:"directory_ldap,omitempty"`

	// The id of the HPC Cache.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below. Changing this forces a new resource to be created.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The ID of the Key Vault Key which should be used to encrypt the data in this HPC Cache.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Specifies the supported Azure Region where the HPC Cache should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A list of IP Addresses where the HPC Cache can be mounted.
	MountAddresses []*string `json:"mountAddresses,omitempty" tf:"mount_addresses,omitempty"`

	// The IPv4 maximum transmission unit configured for the subnet of the HPC Cache. Possible values range from 576 - 1500. Defaults to 1500.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The NTP server IP Address or FQDN for the HPC Cache. Defaults to time.windows.com.
	NtpServer *string `json:"ntpServer,omitempty" tf:"ntp_server,omitempty"`

	// The name of the Resource Group in which to create the HPC Cache. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU of HPC Cache to use. Possible values are (ReadWrite) - Standard_2G, Standard_4G Standard_8G or (ReadOnly) - Standard_L4_5G, Standard_L9G, and Standard_L16G. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The ID of the Subnet for the HPC Cache. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A mapping of tags to assign to the HPC Cache.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HPCCacheParameters struct {

	// Specifies whether the HPC Cache automatically rotates Encryption Key to the latest version.
	// +kubebuilder:validation:Optional
	AutomaticallyRotateKeyToLatestEnabled *bool `json:"automaticallyRotateKeyToLatestEnabled,omitempty" tf:"automatically_rotate_key_to_latest_enabled,omitempty"`

	// The size of the HPC Cache, in GB. Possible values are 3072, 6144, 12288, 21623, 24576, 43246, 49152 and 86491. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CacheSizeInGb *float64 `json:"cacheSizeInGb,omitempty" tf:"cache_size_in_gb,omitempty"`

	// A dns block as defined below.
	// +kubebuilder:validation:Optional
	DNS []DNSParameters `json:"dns,omitempty" tf:"dns,omitempty"`

	// A default_access_policy block as defined below.
	// +kubebuilder:validation:Optional
	DefaultAccessPolicy []DefaultAccessPolicyParameters `json:"defaultAccessPolicy,omitempty" tf:"default_access_policy,omitempty"`

	// A directory_active_directory block as defined below.
	// +kubebuilder:validation:Optional
	DirectoryActiveDirectory []DirectoryActiveDirectoryParameters `json:"directoryActiveDirectory,omitempty" tf:"directory_active_directory,omitempty"`

	// A directory_flat_file block as defined below.
	// +kubebuilder:validation:Optional
	DirectoryFlatFile []DirectoryFlatFileParameters `json:"directoryFlatFile,omitempty" tf:"directory_flat_file,omitempty"`

	// A directory_ldap block as defined below.
	// +kubebuilder:validation:Optional
	DirectoryLdap []DirectoryLdapParameters `json:"directoryLdap,omitempty" tf:"directory_ldap,omitempty"`

	// An identity block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The ID of the Key Vault Key which should be used to encrypt the data in this HPC Cache.
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Specifies the supported Azure Region where the HPC Cache should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The IPv4 maximum transmission unit configured for the subnet of the HPC Cache. Possible values range from 576 - 1500. Defaults to 1500.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The NTP server IP Address or FQDN for the HPC Cache. Defaults to time.windows.com.
	// +kubebuilder:validation:Optional
	NtpServer *string `json:"ntpServer,omitempty" tf:"ntp_server,omitempty"`

	// The name of the Resource Group in which to create the HPC Cache. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU of HPC Cache to use. Possible values are (ReadWrite) - Standard_2G, Standard_4G Standard_8G or (ReadOnly) - Standard_L4_5G, Standard_L9G, and Standard_L16G. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The ID of the Subnet for the HPC Cache. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the HPC Cache.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this HPC Cache. Changing this forces a new resource to be created.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this HPC Cache. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this HPC Cache. Changing this forces a new resource to be created.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this HPC Cache. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this HPC Cache. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this HPC Cache. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// HPCCacheSpec defines the desired state of HPCCache
type HPCCacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HPCCacheParameters `json:"forProvider"`
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
	InitProvider HPCCacheInitParameters `json:"initProvider,omitempty"`
}

// HPCCacheStatus defines the observed state of HPCCache.
type HPCCacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HPCCacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCache is the Schema for the HPCCaches API. Manages a HPC Cache.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HPCCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cacheSizeInGb) || has(self.initProvider.cacheSizeInGb)",message="cacheSizeInGb is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || has(self.initProvider.skuName)",message="skuName is a required parameter"
	Spec   HPCCacheSpec   `json:"spec"`
	Status HPCCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheList contains a list of HPCCaches
type HPCCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HPCCache `json:"items"`
}

// Repository type metadata.
var (
	HPCCache_Kind             = "HPCCache"
	HPCCache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HPCCache_Kind}.String()
	HPCCache_KindAPIVersion   = HPCCache_Kind + "." + CRDGroupVersion.String()
	HPCCache_GroupVersionKind = CRDGroupVersion.WithKind(HPCCache_Kind)
)

func init() {
	SchemeBuilder.Register(&HPCCache{}, &HPCCacheList{})
}
