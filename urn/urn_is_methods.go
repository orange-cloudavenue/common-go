/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package urn

// IsAppPortProfile returns true if the URN is a AppPortProfile URN.
func (urn URN) IsAppPortProfile() bool {
	return urn.IsType(AppPortProfile)
}

// IsOrg returns true if the URN is an Org URN.
func (urn URN) IsOrg() bool {
	return urn.IsType(Org)
}

// IsVM returns true if the URN is a VM URN.
func (urn URN) IsVM() bool {
	return urn.IsType(VM)
}

// IsUser returns true if the URN is a User URN.
func (urn URN) IsUser() bool {
	return urn.IsType(User)
}

// IsGroup returns true if the URN is a Group URN.
func (urn URN) IsGroup() bool {
	return urn.IsType(Group)
}

// IsEdgeGateway returns true if the URN is an EdgeGateway URN.
func (urn URN) IsEdgeGateway() bool {
	return urn.IsType(EdgeGateway)
}

// IsVDC returns true if the URN is a VDC URN.
func (urn URN) IsVDC() bool {
	return urn.IsType(VDC)
}

// IsVDCGroup returns true if the URN is a VDCGroup URN.
func (urn URN) IsVDCGroup() bool {
	return urn.IsType(VDCGroup)
}

// IsNetwork returns true if the URN is a Network URN.
func (urn URN) IsNetwork() bool {
	return urn.IsType(Network)
}

// IsLoadBalancerPool returns true if the URN is a LoadBalancerPool URN.
func (urn URN) IsLoadBalancerPool() bool {
	return urn.IsType(LoadBalancerPool)
}

// IsVDCStorageProfile returns true if the URN is a VDCStorageProfile URN.
func (urn URN) IsVDCStorageProfile() bool {
	return urn.IsType(VDCStorageProfile)
}

// IsVAPP returns true if the URN is a VAPP URN.
func (urn URN) IsVAPP() bool {
	return urn.IsType(VAPP)
}

// IsVAPPTemplate returns true if the URN is a VAPPTemplate URN.
func (urn URN) IsVAPPTemplate() bool {
	return urn.IsType(VAPPTemplate)
}

// IsDisk returns true if the URN is a Disk URN.
func (urn URN) IsDisk() bool {
	return urn.IsType(Disk)
}

// IsSecurityGroup returns true if the URN is a SecurityGroup URN.
func (urn URN) IsSecurityGroup() bool {
	return urn.IsType(SecurityGroup)
}

// IsCatalog returns true if the URN is a Catalog URN.
func (urn URN) IsCatalog() bool {
	return urn.IsType(Catalog)
}

// IsToken returns true if the URN is a Token URN.
func (urn URN) IsToken() bool {
	return urn.IsType(Token)
}

// IsVDCComputePolicy returns true if the URN is a VDCComputePolicy URN.
func (urn URN) IsVDCComputePolicy() bool {
	return urn.IsType(VDCComputePolicy)
}

// IsCertificateLibraryItem returns true if the URN is a CertificateLibraryItem URN.
func (urn URN) IsCertificateLibraryItem() bool {
	return urn.IsType(CertificateLibraryItem)
}

// IsLoadBalancerVirtualService returns true if the URN is a LoadBalancerVirtualService URN.
func (urn URN) IsLoadBalancerVirtualService() bool {
	return urn.IsType(LoadBalancerVirtualService)
}

// IsServiceEngineGroup returns true if the URN is a ServiceEngineGroup URN.
func (urn URN) IsServiceEngineGroup() bool {
	return urn.IsType(ServiceEngineGroup)
}

// IsVCDA returns true if the URN is a VCDA URN.
func (urn URN) IsVCDA() bool {
	return urn.IsType(VCDA)
}
