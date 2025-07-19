# NfProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NfInstanceName** | Pointer to **string** |  | [optional] 
**NfType** | [**NfType**](NfType.md) |  | 
**NfStatus** | [**NfStatus**](NfStatus.md) |  | 
**CollocatedNfInstances** | Pointer to [**[]CollocatedNfInstance**](CollocatedNfInstance.md) |  | [optional] 
**HeartBeatTimer** | Pointer to **int32** |  | [optional] 
**PlmnList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**SnpnList** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**SNssais** | Pointer to [**[]ExtSnssai**](ExtSnssai.md) |  | [optional] 
**PerPlmnSnssaiList** | Pointer to [**[]PlmnSnssai**](PlmnSnssai.md) |  | [optional] 
**NsiList** | Pointer to **[]string** |  | [optional] 
**Fqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**InterPlmnFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**Ipv4Addresses** | Pointer to **[]string** |  | [optional] 
**Ipv6Addresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**AllowedPlmns** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**AllowedSnpns** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**AllowedNfTypes** | Pointer to [**[]NfType**](NfType.md) |  | [optional] 
**AllowedNfDomains** | Pointer to **[]string** |  | [optional] 
**AllowedNssais** | Pointer to [**[]ExtSnssai**](ExtSnssai.md) |  | [optional] 
**AllowedRuleSet** | Pointer to [**map[string]RuleSet**](RuleSet.md) | A map (list of key-value pairs) where a valid JSON pointer Id serves as key | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 
**Load** | Pointer to **int32** |  | [optional] 
**LoadTimeStamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**Locality** | Pointer to **string** |  | [optional] 
**ExtLocality** | Pointer to **map[string]string** | A map (list of key-value pairs) where a (unique) valid JSON string serves as key representing a type of locality  | [optional] 
**UdrInfo** | Pointer to [**UdrInfo**](UdrInfo.md) |  | [optional] 
**UdrInfoList** | Pointer to [**map[string]UdrInfo**](UdrInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdrInfo  | [optional] 
**UdmInfo** | Pointer to [**UdmInfo**](UdmInfo.md) |  | [optional] 
**UdmInfoList** | Pointer to [**map[string]UdmInfo**](UdmInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdmInfo  | [optional] 
**AusfInfo** | Pointer to [**AusfInfo**](AusfInfo.md) |  | [optional] 
**AusfInfoList** | Pointer to [**map[string]AusfInfo**](AusfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AusfInfo  | [optional] 
**AmfInfo** | Pointer to [**AmfInfo**](AmfInfo.md) |  | [optional] 
**AmfInfoList** | Pointer to [**map[string]AmfInfo**](AmfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AmfInfo  | [optional] 
**SmfInfo** | Pointer to [**SmfInfo**](SmfInfo.md) |  | [optional] 
**SmfInfoList** | Pointer to [**map[string]SmfInfo**](SmfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of SmfInfo  | [optional] 
**UpfInfo** | Pointer to [**UpfInfo**](UpfInfo.md) |  | [optional] 
**UpfInfoList** | Pointer to [**map[string]UpfInfo**](UpfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UpfInfo  | [optional] 
**PcfInfo** | Pointer to [**PcfInfo**](PcfInfo.md) |  | [optional] 
**PcfInfoList** | Pointer to [**map[string]PcfInfo**](PcfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcfInfo  | [optional] 
**BsfInfo** | Pointer to [**BsfInfo**](BsfInfo.md) |  | [optional] 
**BsfInfoList** | Pointer to [**map[string]BsfInfo**](BsfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of BsfInfo  | [optional] 
**ChfInfo** | Pointer to [**ChfInfo**](ChfInfo.md) |  | [optional] 
**ChfInfoList** | Pointer to [**map[string]ChfInfo**](ChfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of ChfInfo  | [optional] 
**NefInfo** | Pointer to [**NefInfo**](NefInfo.md) |  | [optional] 
**NrfInfo** | Pointer to [**NrfInfo**](NrfInfo.md) |  | [optional] 
**UdsfInfo** | Pointer to [**UdsfInfo**](UdsfInfo.md) |  | [optional] 
**UdsfInfoList** | Pointer to [**map[string]UdsfInfo**](UdsfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdsfInfo  | [optional] 
**NwdafInfo** | Pointer to [**NwdafInfo**](NwdafInfo.md) |  | [optional] 
**NwdafInfoList** | Pointer to [**map[string]NwdafInfo**](NwdafInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of NwdafInfo  | [optional] 
**PcscfInfoList** | Pointer to [**map[string]PcscfInfo**](PcscfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcscfInfo  | [optional] 
**HssInfoList** | Pointer to [**map[string]HssInfo**](HssInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of HssInfo  | [optional] 
**CustomInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**NfServicePersistence** | Pointer to **bool** |  | [optional] [default to false]
**NfServices** | Pointer to [**[]NFService**](NFService.md) |  | [optional] 
**NfServiceList** | Pointer to [**map[string]NFService**](NFService.md) | A map (list of key-value pairs) where serviceInstanceId serves as key of NFService  | [optional] 
**NfProfileChangesSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**NfProfilePartialUpdateChangesSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**NfProfileChangesInd** | Pointer to **bool** |  | [optional] [readonly] [default to false]
**DefaultNotificationSubscriptions** | Pointer to [**[]DefaultNotificationSubscription**](DefaultNotificationSubscription.md) |  | [optional] 
**LmfInfo** | Pointer to [**LmfInfo**](LmfInfo.md) |  | [optional] 
**GmlcInfo** | Pointer to [**GmlcInfo**](GmlcInfo.md) |  | [optional] 
**NfSetIdList** | Pointer to **[]string** |  | [optional] 
**ServingScope** | Pointer to **[]string** |  | [optional] 
**LcHSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**OlcHSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**NfSetRecoveryTimeList** | Pointer to [**map[string]time.Time**](time.Time.md) | A map (list of key-value pairs) where NfSetId serves as key of DateTime | [optional] 
**ServiceSetRecoveryTimeList** | Pointer to [**map[string]time.Time**](time.Time.md) | A map (list of key-value pairs) where NfServiceSetId serves as key of DateTime  | [optional] 
**ScpDomains** | Pointer to **[]string** |  | [optional] 
**ScpInfo** | Pointer to [**ScpInfo**](ScpInfo.md) |  | [optional] 
**SeppInfo** | Pointer to [**SeppInfo**](SeppInfo.md) |  | [optional] 
**VendorId** | Pointer to **string** | Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA) | [optional] 
**SupportedVendorSpecificFeatures** | Pointer to [**map[string][]VendorSpecificFeature**](array.md) | The key of the map is the IANA-assigned SMI Network Management Private Enterprise Codes  | [optional] 
**AanfInfoList** | Pointer to [**map[string]AanfInfo**](AanfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AanfInfo  | [optional] 
**Var5gDdnmfInfo** | Pointer to [**Model5GDdnmfInfo**](Model5GDdnmfInfo.md) |  | [optional] 
**MfafInfo** | Pointer to [**MfafInfo**](MfafInfo.md) |  | [optional] 
**EasdfInfoList** | Pointer to [**map[string]EasdfInfo**](EasdfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of EasdfInfo  | [optional] 
**DccfInfo** | Pointer to [**DccfInfo**](DccfInfo.md) |  | [optional] 
**NsacfInfoList** | Pointer to [**map[string]NsacfInfo**](NsacfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of NsacfInfo  | [optional] 
**MbSmfInfoList** | Pointer to [**map[string]MbSmfInfo**](MbSmfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbSmfInfo  | [optional] 
**TsctsfInfoList** | Pointer to [**map[string]TsctsfInfo**](TsctsfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of TsctsfInfo  | [optional] 
**MbUpfInfoList** | Pointer to [**map[string]MbUpfInfo**](MbUpfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbUpfInfo  | [optional] 
**TrustAfInfo** | Pointer to [**TrustAfInfo**](TrustAfInfo.md) |  | [optional] 
**NssaafInfo** | Pointer to [**NssaafInfo**](NssaafInfo.md) |  | [optional] 
**HniList** | Pointer to **[]string** |  | [optional] 
**IwmscInfo** | Pointer to [**IwmscInfo**](IwmscInfo.md) |  | [optional] 
**MnpfInfo** | Pointer to [**MnpfInfo**](MnpfInfo.md) |  | [optional] 
**SmsfInfo** | Pointer to [**SmsfInfo**](SmsfInfo.md) |  | [optional] 
**DcsfInfoList** | Pointer to [**map[string]DcsfInfo**](DcsfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of DcsfInfo  | [optional] 
**MrfInfoList** | Pointer to [**map[string]MrfInfo**](MrfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MrfInfo  | [optional] 
**MrfpInfoList** | Pointer to [**map[string]MrfpInfo**](MrfpInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MrfpInfo  | [optional] 
**MfInfoList** | Pointer to [**map[string]MfInfo**](MfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MfInfo  | [optional] 
**AdrfInfoList** | Pointer to [**map[string]AdrfInfo**](AdrfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AdrfInfo  | [optional] 
**SelectionConditions** | Pointer to [**SelectionConditions**](SelectionConditions.md) |  | [optional] 
**CanaryRelease** | Pointer to **bool** |  | [optional] [default to false]
**ExclusiveCanaryReleaseSelection** | Pointer to **bool** |  | [optional] [default to false]
**SharedProfileDataId** | Pointer to **string** |  | [optional] 
**ShutdownTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SupportedRcfs** | Pointer to **[]string** |  | [optional] 
**CanaryPrecedenceOverPreferred** | Pointer to **bool** |  | [optional] [default to false]
**ImsasInfoList** | Pointer to [**map[string]ImsasInfo**](ImsasInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of ImsasInfo  | [optional] 

## Methods

### NewNfProfile

`func NewNfProfile(nfInstanceId string, NfType NfType, NfStatus NfStatus, ) *NfProfile`

NewNfProfile instantiates a new NfProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfProfileWithDefaults

`func NewNfProfileWithDefaults() *NfProfile`

NewNfProfileWithDefaults instantiates a new NfProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *NfProfile) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *NfProfile) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *NfProfile) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetNfInstanceName

`func (o *NfProfile) GetNfInstanceName() string`

GetNfInstanceName returns the NfInstanceName field if non-nil, zero value otherwise.

### GetNfInstanceNameOk

`func (o *NfProfile) GetNfInstanceNameOk() (*string, bool)`

GetNfInstanceNameOk returns a tuple with the NfInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceName

`func (o *NfProfile) SetNfInstanceName(v string)`

SetNfInstanceName sets NfInstanceName field to given value.

### HasNfInstanceName

`func (o *NfProfile) HasNfInstanceName() bool`

HasNfInstanceName returns a boolean if a field has been set.

### GetNfType

`func (o *NfProfile) GetNfType() NfType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *NfProfile) GetNfTypeOk() (*NfType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *NfProfile) SetNfType(v NfType)`

SetNfType sets NfType field to given value.


### GetNfStatus

`func (o *NfProfile) GetNfStatus() NfStatus`

GetNfStatus returns the NfStatus field if non-nil, zero value otherwise.

### GetNfStatusOk

`func (o *NfProfile) GetNfStatusOk() (*NfStatus, bool)`

GetNfStatusOk returns a tuple with the NfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfStatus

`func (o *NfProfile) SetNfStatus(v NfStatus)`

SetNfStatus sets NfStatus field to given value.


### GetCollocatedNfInstances

`func (o *NfProfile) GetCollocatedNfInstances() []CollocatedNfInstance`

GetCollocatedNfInstances returns the CollocatedNfInstances field if non-nil, zero value otherwise.

### GetCollocatedNfInstancesOk

`func (o *NfProfile) GetCollocatedNfInstancesOk() (*[]CollocatedNfInstance, bool)`

GetCollocatedNfInstancesOk returns a tuple with the CollocatedNfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollocatedNfInstances

`func (o *NfProfile) SetCollocatedNfInstances(v []CollocatedNfInstance)`

SetCollocatedNfInstances sets CollocatedNfInstances field to given value.

### HasCollocatedNfInstances

`func (o *NfProfile) HasCollocatedNfInstances() bool`

HasCollocatedNfInstances returns a boolean if a field has been set.

### GetHeartBeatTimer

`func (o *NfProfile) GetHeartBeatTimer() int32`

GetHeartBeatTimer returns the HeartBeatTimer field if non-nil, zero value otherwise.

### GetHeartBeatTimerOk

`func (o *NfProfile) GetHeartBeatTimerOk() (*int32, bool)`

GetHeartBeatTimerOk returns a tuple with the HeartBeatTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartBeatTimer

`func (o *NfProfile) SetHeartBeatTimer(v int32)`

SetHeartBeatTimer sets HeartBeatTimer field to given value.

### HasHeartBeatTimer

`func (o *NfProfile) HasHeartBeatTimer() bool`

HasHeartBeatTimer returns a boolean if a field has been set.

### GetPlmnList

`func (o *NfProfile) GetPlmnList() []PlmnId`

GetPlmnList returns the PlmnList field if non-nil, zero value otherwise.

### GetPlmnListOk

`func (o *NfProfile) GetPlmnListOk() (*[]PlmnId, bool)`

GetPlmnListOk returns a tuple with the PlmnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnList

`func (o *NfProfile) SetPlmnList(v []PlmnId)`

SetPlmnList sets PlmnList field to given value.

### HasPlmnList

`func (o *NfProfile) HasPlmnList() bool`

HasPlmnList returns a boolean if a field has been set.

### GetSnpnList

`func (o *NfProfile) GetSnpnList() []PlmnIdNid`

GetSnpnList returns the SnpnList field if non-nil, zero value otherwise.

### GetSnpnListOk

`func (o *NfProfile) GetSnpnListOk() (*[]PlmnIdNid, bool)`

GetSnpnListOk returns a tuple with the SnpnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnpnList

`func (o *NfProfile) SetSnpnList(v []PlmnIdNid)`

SetSnpnList sets SnpnList field to given value.

### HasSnpnList

`func (o *NfProfile) HasSnpnList() bool`

HasSnpnList returns a boolean if a field has been set.

### GetSNssais

`func (o *NfProfile) GetSNssais() []ExtSnssai`

GetSNssais returns the SNssais field if non-nil, zero value otherwise.

### GetSNssaisOk

`func (o *NfProfile) GetSNssaisOk() (*[]ExtSnssai, bool)`

GetSNssaisOk returns a tuple with the SNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssais

`func (o *NfProfile) SetSNssais(v []ExtSnssai)`

SetSNssais sets SNssais field to given value.

### HasSNssais

`func (o *NfProfile) HasSNssais() bool`

HasSNssais returns a boolean if a field has been set.

### GetPerPlmnSnssaiList

`func (o *NfProfile) GetPerPlmnSnssaiList() []PlmnSnssai`

GetPerPlmnSnssaiList returns the PerPlmnSnssaiList field if non-nil, zero value otherwise.

### GetPerPlmnSnssaiListOk

`func (o *NfProfile) GetPerPlmnSnssaiListOk() (*[]PlmnSnssai, bool)`

GetPerPlmnSnssaiListOk returns a tuple with the PerPlmnSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPlmnSnssaiList

`func (o *NfProfile) SetPerPlmnSnssaiList(v []PlmnSnssai)`

SetPerPlmnSnssaiList sets PerPlmnSnssaiList field to given value.

### HasPerPlmnSnssaiList

`func (o *NfProfile) HasPerPlmnSnssaiList() bool`

HasPerPlmnSnssaiList returns a boolean if a field has been set.

### GetNsiList

`func (o *NfProfile) GetNsiList() []string`

GetNsiList returns the NsiList field if non-nil, zero value otherwise.

### GetNsiListOk

`func (o *NfProfile) GetNsiListOk() (*[]string, bool)`

GetNsiListOk returns a tuple with the NsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiList

`func (o *NfProfile) SetNsiList(v []string)`

SetNsiList sets NsiList field to given value.

### HasNsiList

`func (o *NfProfile) HasNsiList() bool`

HasNsiList returns a boolean if a field has been set.

### GetFqdn

`func (o *NfProfile) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NfProfile) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NfProfile) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NfProfile) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetInterPlmnFqdn

`func (o *NfProfile) GetInterPlmnFqdn() string`

GetInterPlmnFqdn returns the InterPlmnFqdn field if non-nil, zero value otherwise.

### GetInterPlmnFqdnOk

`func (o *NfProfile) GetInterPlmnFqdnOk() (*string, bool)`

GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnFqdn

`func (o *NfProfile) SetInterPlmnFqdn(v string)`

SetInterPlmnFqdn sets InterPlmnFqdn field to given value.

### HasInterPlmnFqdn

`func (o *NfProfile) HasInterPlmnFqdn() bool`

HasInterPlmnFqdn returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *NfProfile) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *NfProfile) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *NfProfile) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *NfProfile) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *NfProfile) GetIpv6Addresses() []Ipv6Addr`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *NfProfile) GetIpv6AddressesOk() (*[]Ipv6Addr, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *NfProfile) SetIpv6Addresses(v []Ipv6Addr)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *NfProfile) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetAllowedPlmns

`func (o *NfProfile) GetAllowedPlmns() []PlmnId`

GetAllowedPlmns returns the AllowedPlmns field if non-nil, zero value otherwise.

### GetAllowedPlmnsOk

`func (o *NfProfile) GetAllowedPlmnsOk() (*[]PlmnId, bool)`

GetAllowedPlmnsOk returns a tuple with the AllowedPlmns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPlmns

`func (o *NfProfile) SetAllowedPlmns(v []PlmnId)`

SetAllowedPlmns sets AllowedPlmns field to given value.

### HasAllowedPlmns

`func (o *NfProfile) HasAllowedPlmns() bool`

HasAllowedPlmns returns a boolean if a field has been set.

### GetAllowedSnpns

`func (o *NfProfile) GetAllowedSnpns() []PlmnIdNid`

GetAllowedSnpns returns the AllowedSnpns field if non-nil, zero value otherwise.

### GetAllowedSnpnsOk

`func (o *NfProfile) GetAllowedSnpnsOk() (*[]PlmnIdNid, bool)`

GetAllowedSnpnsOk returns a tuple with the AllowedSnpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSnpns

`func (o *NfProfile) SetAllowedSnpns(v []PlmnIdNid)`

SetAllowedSnpns sets AllowedSnpns field to given value.

### HasAllowedSnpns

`func (o *NfProfile) HasAllowedSnpns() bool`

HasAllowedSnpns returns a boolean if a field has been set.

### GetAllowedNfTypes

`func (o *NfProfile) GetAllowedNfTypes() []NfType`

GetAllowedNfTypes returns the AllowedNfTypes field if non-nil, zero value otherwise.

### GetAllowedNfTypesOk

`func (o *NfProfile) GetAllowedNfTypesOk() (*[]NfType, bool)`

GetAllowedNfTypesOk returns a tuple with the AllowedNfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfTypes

`func (o *NfProfile) SetAllowedNfTypes(v []NfType)`

SetAllowedNfTypes sets AllowedNfTypes field to given value.

### HasAllowedNfTypes

`func (o *NfProfile) HasAllowedNfTypes() bool`

HasAllowedNfTypes returns a boolean if a field has been set.

### GetAllowedNfDomains

`func (o *NfProfile) GetAllowedNfDomains() []string`

GetAllowedNfDomains returns the AllowedNfDomains field if non-nil, zero value otherwise.

### GetAllowedNfDomainsOk

`func (o *NfProfile) GetAllowedNfDomainsOk() (*[]string, bool)`

GetAllowedNfDomainsOk returns a tuple with the AllowedNfDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfDomains

`func (o *NfProfile) SetAllowedNfDomains(v []string)`

SetAllowedNfDomains sets AllowedNfDomains field to given value.

### HasAllowedNfDomains

`func (o *NfProfile) HasAllowedNfDomains() bool`

HasAllowedNfDomains returns a boolean if a field has been set.

### GetAllowedNssais

`func (o *NfProfile) GetAllowedNssais() []ExtSnssai`

GetAllowedNssais returns the AllowedNssais field if non-nil, zero value otherwise.

### GetAllowedNssaisOk

`func (o *NfProfile) GetAllowedNssaisOk() (*[]ExtSnssai, bool)`

GetAllowedNssaisOk returns a tuple with the AllowedNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssais

`func (o *NfProfile) SetAllowedNssais(v []ExtSnssai)`

SetAllowedNssais sets AllowedNssais field to given value.

### HasAllowedNssais

`func (o *NfProfile) HasAllowedNssais() bool`

HasAllowedNssais returns a boolean if a field has been set.

### GetAllowedRuleSet

`func (o *NfProfile) GetAllowedRuleSet() map[string]RuleSet`

GetAllowedRuleSet returns the AllowedRuleSet field if non-nil, zero value otherwise.

### GetAllowedRuleSetOk

`func (o *NfProfile) GetAllowedRuleSetOk() (*map[string]RuleSet, bool)`

GetAllowedRuleSetOk returns a tuple with the AllowedRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRuleSet

`func (o *NfProfile) SetAllowedRuleSet(v map[string]RuleSet)`

SetAllowedRuleSet sets AllowedRuleSet field to given value.

### HasAllowedRuleSet

`func (o *NfProfile) HasAllowedRuleSet() bool`

HasAllowedRuleSet returns a boolean if a field has been set.

### GetPriority

`func (o *NfProfile) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NfProfile) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NfProfile) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NfProfile) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCapacity

`func (o *NfProfile) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *NfProfile) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *NfProfile) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *NfProfile) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetLoad

`func (o *NfProfile) GetLoad() int32`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *NfProfile) GetLoadOk() (*int32, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *NfProfile) SetLoad(v int32)`

SetLoad sets Load field to given value.

### HasLoad

`func (o *NfProfile) HasLoad() bool`

HasLoad returns a boolean if a field has been set.

### GetLoadTimeStamp

`func (o *NfProfile) GetLoadTimeStamp() time.Time`

GetLoadTimeStamp returns the LoadTimeStamp field if non-nil, zero value otherwise.

### GetLoadTimeStampOk

`func (o *NfProfile) GetLoadTimeStampOk() (*time.Time, bool)`

GetLoadTimeStampOk returns a tuple with the LoadTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTimeStamp

`func (o *NfProfile) SetLoadTimeStamp(v time.Time)`

SetLoadTimeStamp sets LoadTimeStamp field to given value.

### HasLoadTimeStamp

`func (o *NfProfile) HasLoadTimeStamp() bool`

HasLoadTimeStamp returns a boolean if a field has been set.

### GetLocality

`func (o *NfProfile) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *NfProfile) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *NfProfile) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *NfProfile) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetExtLocality

`func (o *NfProfile) GetExtLocality() map[string]string`

GetExtLocality returns the ExtLocality field if non-nil, zero value otherwise.

### GetExtLocalityOk

`func (o *NfProfile) GetExtLocalityOk() (*map[string]string, bool)`

GetExtLocalityOk returns a tuple with the ExtLocality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtLocality

`func (o *NfProfile) SetExtLocality(v map[string]string)`

SetExtLocality sets ExtLocality field to given value.

### HasExtLocality

`func (o *NfProfile) HasExtLocality() bool`

HasExtLocality returns a boolean if a field has been set.

### GetUdrInfo

`func (o *NfProfile) GetUdrInfo() UdrInfo`

GetUdrInfo returns the UdrInfo field if non-nil, zero value otherwise.

### GetUdrInfoOk

`func (o *NfProfile) GetUdrInfoOk() (*UdrInfo, bool)`

GetUdrInfoOk returns a tuple with the UdrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrInfo

`func (o *NfProfile) SetUdrInfo(v UdrInfo)`

SetUdrInfo sets UdrInfo field to given value.

### HasUdrInfo

`func (o *NfProfile) HasUdrInfo() bool`

HasUdrInfo returns a boolean if a field has been set.

### GetUdrInfoList

`func (o *NfProfile) GetUdrInfoList() map[string]UdrInfo`

GetUdrInfoList returns the UdrInfoList field if non-nil, zero value otherwise.

### GetUdrInfoListOk

`func (o *NfProfile) GetUdrInfoListOk() (*map[string]UdrInfo, bool)`

GetUdrInfoListOk returns a tuple with the UdrInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrInfoList

`func (o *NfProfile) SetUdrInfoList(v map[string]UdrInfo)`

SetUdrInfoList sets UdrInfoList field to given value.

### HasUdrInfoList

`func (o *NfProfile) HasUdrInfoList() bool`

HasUdrInfoList returns a boolean if a field has been set.

### GetUdmInfo

`func (o *NfProfile) GetUdmInfo() UdmInfo`

GetUdmInfo returns the UdmInfo field if non-nil, zero value otherwise.

### GetUdmInfoOk

`func (o *NfProfile) GetUdmInfoOk() (*UdmInfo, bool)`

GetUdmInfoOk returns a tuple with the UdmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmInfo

`func (o *NfProfile) SetUdmInfo(v UdmInfo)`

SetUdmInfo sets UdmInfo field to given value.

### HasUdmInfo

`func (o *NfProfile) HasUdmInfo() bool`

HasUdmInfo returns a boolean if a field has been set.

### GetUdmInfoList

`func (o *NfProfile) GetUdmInfoList() map[string]UdmInfo`

GetUdmInfoList returns the UdmInfoList field if non-nil, zero value otherwise.

### GetUdmInfoListOk

`func (o *NfProfile) GetUdmInfoListOk() (*map[string]UdmInfo, bool)`

GetUdmInfoListOk returns a tuple with the UdmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmInfoList

`func (o *NfProfile) SetUdmInfoList(v map[string]UdmInfo)`

SetUdmInfoList sets UdmInfoList field to given value.

### HasUdmInfoList

`func (o *NfProfile) HasUdmInfoList() bool`

HasUdmInfoList returns a boolean if a field has been set.

### GetAusfInfo

`func (o *NfProfile) GetAusfInfo() AusfInfo`

GetAusfInfo returns the AusfInfo field if non-nil, zero value otherwise.

### GetAusfInfoOk

`func (o *NfProfile) GetAusfInfoOk() (*AusfInfo, bool)`

GetAusfInfoOk returns a tuple with the AusfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfInfo

`func (o *NfProfile) SetAusfInfo(v AusfInfo)`

SetAusfInfo sets AusfInfo field to given value.

### HasAusfInfo

`func (o *NfProfile) HasAusfInfo() bool`

HasAusfInfo returns a boolean if a field has been set.

### GetAusfInfoList

`func (o *NfProfile) GetAusfInfoList() map[string]AusfInfo`

GetAusfInfoList returns the AusfInfoList field if non-nil, zero value otherwise.

### GetAusfInfoListOk

`func (o *NfProfile) GetAusfInfoListOk() (*map[string]AusfInfo, bool)`

GetAusfInfoListOk returns a tuple with the AusfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfInfoList

`func (o *NfProfile) SetAusfInfoList(v map[string]AusfInfo)`

SetAusfInfoList sets AusfInfoList field to given value.

### HasAusfInfoList

`func (o *NfProfile) HasAusfInfoList() bool`

HasAusfInfoList returns a boolean if a field has been set.

### GetAmfInfo

`func (o *NfProfile) GetAmfInfo() AmfInfo`

GetAmfInfo returns the AmfInfo field if non-nil, zero value otherwise.

### GetAmfInfoOk

`func (o *NfProfile) GetAmfInfoOk() (*AmfInfo, bool)`

GetAmfInfoOk returns a tuple with the AmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInfo

`func (o *NfProfile) SetAmfInfo(v AmfInfo)`

SetAmfInfo sets AmfInfo field to given value.

### HasAmfInfo

`func (o *NfProfile) HasAmfInfo() bool`

HasAmfInfo returns a boolean if a field has been set.

### GetAmfInfoList

`func (o *NfProfile) GetAmfInfoList() map[string]AmfInfo`

GetAmfInfoList returns the AmfInfoList field if non-nil, zero value otherwise.

### GetAmfInfoListOk

`func (o *NfProfile) GetAmfInfoListOk() (*map[string]AmfInfo, bool)`

GetAmfInfoListOk returns a tuple with the AmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInfoList

`func (o *NfProfile) SetAmfInfoList(v map[string]AmfInfo)`

SetAmfInfoList sets AmfInfoList field to given value.

### HasAmfInfoList

`func (o *NfProfile) HasAmfInfoList() bool`

HasAmfInfoList returns a boolean if a field has been set.

### GetSmfInfo

`func (o *NfProfile) GetSmfInfo() SmfInfo`

GetSmfInfo returns the SmfInfo field if non-nil, zero value otherwise.

### GetSmfInfoOk

`func (o *NfProfile) GetSmfInfoOk() (*SmfInfo, bool)`

GetSmfInfoOk returns a tuple with the SmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInfo

`func (o *NfProfile) SetSmfInfo(v SmfInfo)`

SetSmfInfo sets SmfInfo field to given value.

### HasSmfInfo

`func (o *NfProfile) HasSmfInfo() bool`

HasSmfInfo returns a boolean if a field has been set.

### GetSmfInfoList

`func (o *NfProfile) GetSmfInfoList() map[string]SmfInfo`

GetSmfInfoList returns the SmfInfoList field if non-nil, zero value otherwise.

### GetSmfInfoListOk

`func (o *NfProfile) GetSmfInfoListOk() (*map[string]SmfInfo, bool)`

GetSmfInfoListOk returns a tuple with the SmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInfoList

`func (o *NfProfile) SetSmfInfoList(v map[string]SmfInfo)`

SetSmfInfoList sets SmfInfoList field to given value.

### HasSmfInfoList

`func (o *NfProfile) HasSmfInfoList() bool`

HasSmfInfoList returns a boolean if a field has been set.

### GetUpfInfo

`func (o *NfProfile) GetUpfInfo() UpfInfo`

GetUpfInfo returns the UpfInfo field if non-nil, zero value otherwise.

### GetUpfInfoOk

`func (o *NfProfile) GetUpfInfoOk() (*UpfInfo, bool)`

GetUpfInfoOk returns a tuple with the UpfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfInfo

`func (o *NfProfile) SetUpfInfo(v UpfInfo)`

SetUpfInfo sets UpfInfo field to given value.

### HasUpfInfo

`func (o *NfProfile) HasUpfInfo() bool`

HasUpfInfo returns a boolean if a field has been set.

### GetUpfInfoList

`func (o *NfProfile) GetUpfInfoList() map[string]UpfInfo`

GetUpfInfoList returns the UpfInfoList field if non-nil, zero value otherwise.

### GetUpfInfoListOk

`func (o *NfProfile) GetUpfInfoListOk() (*map[string]UpfInfo, bool)`

GetUpfInfoListOk returns a tuple with the UpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfInfoList

`func (o *NfProfile) SetUpfInfoList(v map[string]UpfInfo)`

SetUpfInfoList sets UpfInfoList field to given value.

### HasUpfInfoList

`func (o *NfProfile) HasUpfInfoList() bool`

HasUpfInfoList returns a boolean if a field has been set.

### GetPcfInfo

`func (o *NfProfile) GetPcfInfo() PcfInfo`

GetPcfInfo returns the PcfInfo field if non-nil, zero value otherwise.

### GetPcfInfoOk

`func (o *NfProfile) GetPcfInfoOk() (*PcfInfo, bool)`

GetPcfInfoOk returns a tuple with the PcfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfInfo

`func (o *NfProfile) SetPcfInfo(v PcfInfo)`

SetPcfInfo sets PcfInfo field to given value.

### HasPcfInfo

`func (o *NfProfile) HasPcfInfo() bool`

HasPcfInfo returns a boolean if a field has been set.

### GetPcfInfoList

`func (o *NfProfile) GetPcfInfoList() map[string]PcfInfo`

GetPcfInfoList returns the PcfInfoList field if non-nil, zero value otherwise.

### GetPcfInfoListOk

`func (o *NfProfile) GetPcfInfoListOk() (*map[string]PcfInfo, bool)`

GetPcfInfoListOk returns a tuple with the PcfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfInfoList

`func (o *NfProfile) SetPcfInfoList(v map[string]PcfInfo)`

SetPcfInfoList sets PcfInfoList field to given value.

### HasPcfInfoList

`func (o *NfProfile) HasPcfInfoList() bool`

HasPcfInfoList returns a boolean if a field has been set.

### GetBsfInfo

`func (o *NfProfile) GetBsfInfo() BsfInfo`

GetBsfInfo returns the BsfInfo field if non-nil, zero value otherwise.

### GetBsfInfoOk

`func (o *NfProfile) GetBsfInfoOk() (*BsfInfo, bool)`

GetBsfInfoOk returns a tuple with the BsfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsfInfo

`func (o *NfProfile) SetBsfInfo(v BsfInfo)`

SetBsfInfo sets BsfInfo field to given value.

### HasBsfInfo

`func (o *NfProfile) HasBsfInfo() bool`

HasBsfInfo returns a boolean if a field has been set.

### GetBsfInfoList

`func (o *NfProfile) GetBsfInfoList() map[string]BsfInfo`

GetBsfInfoList returns the BsfInfoList field if non-nil, zero value otherwise.

### GetBsfInfoListOk

`func (o *NfProfile) GetBsfInfoListOk() (*map[string]BsfInfo, bool)`

GetBsfInfoListOk returns a tuple with the BsfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsfInfoList

`func (o *NfProfile) SetBsfInfoList(v map[string]BsfInfo)`

SetBsfInfoList sets BsfInfoList field to given value.

### HasBsfInfoList

`func (o *NfProfile) HasBsfInfoList() bool`

HasBsfInfoList returns a boolean if a field has been set.

### GetChfInfo

`func (o *NfProfile) GetChfInfo() ChfInfo`

GetChfInfo returns the ChfInfo field if non-nil, zero value otherwise.

### GetChfInfoOk

`func (o *NfProfile) GetChfInfoOk() (*ChfInfo, bool)`

GetChfInfoOk returns a tuple with the ChfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChfInfo

`func (o *NfProfile) SetChfInfo(v ChfInfo)`

SetChfInfo sets ChfInfo field to given value.

### HasChfInfo

`func (o *NfProfile) HasChfInfo() bool`

HasChfInfo returns a boolean if a field has been set.

### GetChfInfoList

`func (o *NfProfile) GetChfInfoList() map[string]ChfInfo`

GetChfInfoList returns the ChfInfoList field if non-nil, zero value otherwise.

### GetChfInfoListOk

`func (o *NfProfile) GetChfInfoListOk() (*map[string]ChfInfo, bool)`

GetChfInfoListOk returns a tuple with the ChfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChfInfoList

`func (o *NfProfile) SetChfInfoList(v map[string]ChfInfo)`

SetChfInfoList sets ChfInfoList field to given value.

### HasChfInfoList

`func (o *NfProfile) HasChfInfoList() bool`

HasChfInfoList returns a boolean if a field has been set.

### GetNefInfo

`func (o *NfProfile) GetNefInfo() NefInfo`

GetNefInfo returns the NefInfo field if non-nil, zero value otherwise.

### GetNefInfoOk

`func (o *NfProfile) GetNefInfoOk() (*NefInfo, bool)`

GetNefInfoOk returns a tuple with the NefInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefInfo

`func (o *NfProfile) SetNefInfo(v NefInfo)`

SetNefInfo sets NefInfo field to given value.

### HasNefInfo

`func (o *NfProfile) HasNefInfo() bool`

HasNefInfo returns a boolean if a field has been set.

### GetNrfInfo

`func (o *NfProfile) GetNrfInfo() NrfInfo`

GetNrfInfo returns the NrfInfo field if non-nil, zero value otherwise.

### GetNrfInfoOk

`func (o *NfProfile) GetNrfInfoOk() (*NrfInfo, bool)`

GetNrfInfoOk returns a tuple with the NrfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfInfo

`func (o *NfProfile) SetNrfInfo(v NrfInfo)`

SetNrfInfo sets NrfInfo field to given value.

### HasNrfInfo

`func (o *NfProfile) HasNrfInfo() bool`

HasNrfInfo returns a boolean if a field has been set.

### GetUdsfInfo

`func (o *NfProfile) GetUdsfInfo() UdsfInfo`

GetUdsfInfo returns the UdsfInfo field if non-nil, zero value otherwise.

### GetUdsfInfoOk

`func (o *NfProfile) GetUdsfInfoOk() (*UdsfInfo, bool)`

GetUdsfInfoOk returns a tuple with the UdsfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsfInfo

`func (o *NfProfile) SetUdsfInfo(v UdsfInfo)`

SetUdsfInfo sets UdsfInfo field to given value.

### HasUdsfInfo

`func (o *NfProfile) HasUdsfInfo() bool`

HasUdsfInfo returns a boolean if a field has been set.

### GetUdsfInfoList

`func (o *NfProfile) GetUdsfInfoList() map[string]UdsfInfo`

GetUdsfInfoList returns the UdsfInfoList field if non-nil, zero value otherwise.

### GetUdsfInfoListOk

`func (o *NfProfile) GetUdsfInfoListOk() (*map[string]UdsfInfo, bool)`

GetUdsfInfoListOk returns a tuple with the UdsfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsfInfoList

`func (o *NfProfile) SetUdsfInfoList(v map[string]UdsfInfo)`

SetUdsfInfoList sets UdsfInfoList field to given value.

### HasUdsfInfoList

`func (o *NfProfile) HasUdsfInfoList() bool`

HasUdsfInfoList returns a boolean if a field has been set.

### GetNwdafInfo

`func (o *NfProfile) GetNwdafInfo() NwdafInfo`

GetNwdafInfo returns the NwdafInfo field if non-nil, zero value otherwise.

### GetNwdafInfoOk

`func (o *NfProfile) GetNwdafInfoOk() (*NwdafInfo, bool)`

GetNwdafInfoOk returns a tuple with the NwdafInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafInfo

`func (o *NfProfile) SetNwdafInfo(v NwdafInfo)`

SetNwdafInfo sets NwdafInfo field to given value.

### HasNwdafInfo

`func (o *NfProfile) HasNwdafInfo() bool`

HasNwdafInfo returns a boolean if a field has been set.

### GetNwdafInfoList

`func (o *NfProfile) GetNwdafInfoList() map[string]NwdafInfo`

GetNwdafInfoList returns the NwdafInfoList field if non-nil, zero value otherwise.

### GetNwdafInfoListOk

`func (o *NfProfile) GetNwdafInfoListOk() (*map[string]NwdafInfo, bool)`

GetNwdafInfoListOk returns a tuple with the NwdafInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafInfoList

`func (o *NfProfile) SetNwdafInfoList(v map[string]NwdafInfo)`

SetNwdafInfoList sets NwdafInfoList field to given value.

### HasNwdafInfoList

`func (o *NfProfile) HasNwdafInfoList() bool`

HasNwdafInfoList returns a boolean if a field has been set.

### GetPcscfInfoList

`func (o *NfProfile) GetPcscfInfoList() map[string]PcscfInfo`

GetPcscfInfoList returns the PcscfInfoList field if non-nil, zero value otherwise.

### GetPcscfInfoListOk

`func (o *NfProfile) GetPcscfInfoListOk() (*map[string]PcscfInfo, bool)`

GetPcscfInfoListOk returns a tuple with the PcscfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfInfoList

`func (o *NfProfile) SetPcscfInfoList(v map[string]PcscfInfo)`

SetPcscfInfoList sets PcscfInfoList field to given value.

### HasPcscfInfoList

`func (o *NfProfile) HasPcscfInfoList() bool`

HasPcscfInfoList returns a boolean if a field has been set.

### GetHssInfoList

`func (o *NfProfile) GetHssInfoList() map[string]HssInfo`

GetHssInfoList returns the HssInfoList field if non-nil, zero value otherwise.

### GetHssInfoListOk

`func (o *NfProfile) GetHssInfoListOk() (*map[string]HssInfo, bool)`

GetHssInfoListOk returns a tuple with the HssInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssInfoList

`func (o *NfProfile) SetHssInfoList(v map[string]HssInfo)`

SetHssInfoList sets HssInfoList field to given value.

### HasHssInfoList

`func (o *NfProfile) HasHssInfoList() bool`

HasHssInfoList returns a boolean if a field has been set.

### GetCustomInfo

`func (o *NfProfile) GetCustomInfo() map[string]interface{}`

GetCustomInfo returns the CustomInfo field if non-nil, zero value otherwise.

### GetCustomInfoOk

`func (o *NfProfile) GetCustomInfoOk() (*map[string]interface{}, bool)`

GetCustomInfoOk returns a tuple with the CustomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInfo

`func (o *NfProfile) SetCustomInfo(v map[string]interface{})`

SetCustomInfo sets CustomInfo field to given value.

### HasCustomInfo

`func (o *NfProfile) HasCustomInfo() bool`

HasCustomInfo returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *NfProfile) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *NfProfile) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *NfProfile) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *NfProfile) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetNfServicePersistence

`func (o *NfProfile) GetNfServicePersistence() bool`

GetNfServicePersistence returns the NfServicePersistence field if non-nil, zero value otherwise.

### GetNfServicePersistenceOk

`func (o *NfProfile) GetNfServicePersistenceOk() (*bool, bool)`

GetNfServicePersistenceOk returns a tuple with the NfServicePersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServicePersistence

`func (o *NfProfile) SetNfServicePersistence(v bool)`

SetNfServicePersistence sets NfServicePersistence field to given value.

### HasNfServicePersistence

`func (o *NfProfile) HasNfServicePersistence() bool`

HasNfServicePersistence returns a boolean if a field has been set.

### GetNfServices

`func (o *NfProfile) GetNfServices() []NFService`

GetNfServices returns the NfServices field if non-nil, zero value otherwise.

### GetNfServicesOk

`func (o *NfProfile) GetNfServicesOk() (*[]NFService, bool)`

GetNfServicesOk returns a tuple with the NfServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServices

`func (o *NfProfile) SetNfServices(v []NFService)`

SetNfServices sets NfServices field to given value.

### HasNfServices

`func (o *NfProfile) HasNfServices() bool`

HasNfServices returns a boolean if a field has been set.

### GetNfServiceList

`func (o *NfProfile) GetNfServiceList() map[string]NFService`

GetNfServiceList returns the NfServiceList field if non-nil, zero value otherwise.

### GetNfServiceListOk

`func (o *NfProfile) GetNfServiceListOk() (*map[string]NFService, bool)`

GetNfServiceListOk returns a tuple with the NfServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServiceList

`func (o *NfProfile) SetNfServiceList(v map[string]NFService)`

SetNfServiceList sets NfServiceList field to given value.

### HasNfServiceList

`func (o *NfProfile) HasNfServiceList() bool`

HasNfServiceList returns a boolean if a field has been set.

### GetNfProfileChangesSupportInd

`func (o *NfProfile) GetNfProfileChangesSupportInd() bool`

GetNfProfileChangesSupportInd returns the NfProfileChangesSupportInd field if non-nil, zero value otherwise.

### GetNfProfileChangesSupportIndOk

`func (o *NfProfile) GetNfProfileChangesSupportIndOk() (*bool, bool)`

GetNfProfileChangesSupportIndOk returns a tuple with the NfProfileChangesSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfProfileChangesSupportInd

`func (o *NfProfile) SetNfProfileChangesSupportInd(v bool)`

SetNfProfileChangesSupportInd sets NfProfileChangesSupportInd field to given value.

### HasNfProfileChangesSupportInd

`func (o *NfProfile) HasNfProfileChangesSupportInd() bool`

HasNfProfileChangesSupportInd returns a boolean if a field has been set.

### GetNfProfilePartialUpdateChangesSupportInd

`func (o *NfProfile) GetNfProfilePartialUpdateChangesSupportInd() bool`

GetNfProfilePartialUpdateChangesSupportInd returns the NfProfilePartialUpdateChangesSupportInd field if non-nil, zero value otherwise.

### GetNfProfilePartialUpdateChangesSupportIndOk

`func (o *NfProfile) GetNfProfilePartialUpdateChangesSupportIndOk() (*bool, bool)`

GetNfProfilePartialUpdateChangesSupportIndOk returns a tuple with the NfProfilePartialUpdateChangesSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfProfilePartialUpdateChangesSupportInd

`func (o *NfProfile) SetNfProfilePartialUpdateChangesSupportInd(v bool)`

SetNfProfilePartialUpdateChangesSupportInd sets NfProfilePartialUpdateChangesSupportInd field to given value.

### HasNfProfilePartialUpdateChangesSupportInd

`func (o *NfProfile) HasNfProfilePartialUpdateChangesSupportInd() bool`

HasNfProfilePartialUpdateChangesSupportInd returns a boolean if a field has been set.

### GetNfProfileChangesInd

`func (o *NfProfile) GetNfProfileChangesInd() bool`

GetNfProfileChangesInd returns the NfProfileChangesInd field if non-nil, zero value otherwise.

### GetNfProfileChangesIndOk

`func (o *NfProfile) GetNfProfileChangesIndOk() (*bool, bool)`

GetNfProfileChangesIndOk returns a tuple with the NfProfileChangesInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfProfileChangesInd

`func (o *NfProfile) SetNfProfileChangesInd(v bool)`

SetNfProfileChangesInd sets NfProfileChangesInd field to given value.

### HasNfProfileChangesInd

`func (o *NfProfile) HasNfProfileChangesInd() bool`

HasNfProfileChangesInd returns a boolean if a field has been set.

### GetDefaultNotificationSubscriptions

`func (o *NfProfile) GetDefaultNotificationSubscriptions() []DefaultNotificationSubscription`

GetDefaultNotificationSubscriptions returns the DefaultNotificationSubscriptions field if non-nil, zero value otherwise.

### GetDefaultNotificationSubscriptionsOk

`func (o *NfProfile) GetDefaultNotificationSubscriptionsOk() (*[]DefaultNotificationSubscription, bool)`

GetDefaultNotificationSubscriptionsOk returns a tuple with the DefaultNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNotificationSubscriptions

`func (o *NfProfile) SetDefaultNotificationSubscriptions(v []DefaultNotificationSubscription)`

SetDefaultNotificationSubscriptions sets DefaultNotificationSubscriptions field to given value.

### HasDefaultNotificationSubscriptions

`func (o *NfProfile) HasDefaultNotificationSubscriptions() bool`

HasDefaultNotificationSubscriptions returns a boolean if a field has been set.

### GetLmfInfo

`func (o *NfProfile) GetLmfInfo() LmfInfo`

GetLmfInfo returns the LmfInfo field if non-nil, zero value otherwise.

### GetLmfInfoOk

`func (o *NfProfile) GetLmfInfoOk() (*LmfInfo, bool)`

GetLmfInfoOk returns a tuple with the LmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfInfo

`func (o *NfProfile) SetLmfInfo(v LmfInfo)`

SetLmfInfo sets LmfInfo field to given value.

### HasLmfInfo

`func (o *NfProfile) HasLmfInfo() bool`

HasLmfInfo returns a boolean if a field has been set.

### GetGmlcInfo

`func (o *NfProfile) GetGmlcInfo() GmlcInfo`

GetGmlcInfo returns the GmlcInfo field if non-nil, zero value otherwise.

### GetGmlcInfoOk

`func (o *NfProfile) GetGmlcInfoOk() (*GmlcInfo, bool)`

GetGmlcInfoOk returns a tuple with the GmlcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmlcInfo

`func (o *NfProfile) SetGmlcInfo(v GmlcInfo)`

SetGmlcInfo sets GmlcInfo field to given value.

### HasGmlcInfo

`func (o *NfProfile) HasGmlcInfo() bool`

HasGmlcInfo returns a boolean if a field has been set.

### GetNfSetIdList

`func (o *NfProfile) GetNfSetIdList() []string`

GetNfSetIdList returns the NfSetIdList field if non-nil, zero value otherwise.

### GetNfSetIdListOk

`func (o *NfProfile) GetNfSetIdListOk() (*[]string, bool)`

GetNfSetIdListOk returns a tuple with the NfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetIdList

`func (o *NfProfile) SetNfSetIdList(v []string)`

SetNfSetIdList sets NfSetIdList field to given value.

### HasNfSetIdList

`func (o *NfProfile) HasNfSetIdList() bool`

HasNfSetIdList returns a boolean if a field has been set.

### GetServingScope

`func (o *NfProfile) GetServingScope() []string`

GetServingScope returns the ServingScope field if non-nil, zero value otherwise.

### GetServingScopeOk

`func (o *NfProfile) GetServingScopeOk() (*[]string, bool)`

GetServingScopeOk returns a tuple with the ServingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingScope

`func (o *NfProfile) SetServingScope(v []string)`

SetServingScope sets ServingScope field to given value.

### HasServingScope

`func (o *NfProfile) HasServingScope() bool`

HasServingScope returns a boolean if a field has been set.

### GetLcHSupportInd

`func (o *NfProfile) GetLcHSupportInd() bool`

GetLcHSupportInd returns the LcHSupportInd field if non-nil, zero value otherwise.

### GetLcHSupportIndOk

`func (o *NfProfile) GetLcHSupportIndOk() (*bool, bool)`

GetLcHSupportIndOk returns a tuple with the LcHSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcHSupportInd

`func (o *NfProfile) SetLcHSupportInd(v bool)`

SetLcHSupportInd sets LcHSupportInd field to given value.

### HasLcHSupportInd

`func (o *NfProfile) HasLcHSupportInd() bool`

HasLcHSupportInd returns a boolean if a field has been set.

### GetOlcHSupportInd

`func (o *NfProfile) GetOlcHSupportInd() bool`

GetOlcHSupportInd returns the OlcHSupportInd field if non-nil, zero value otherwise.

### GetOlcHSupportIndOk

`func (o *NfProfile) GetOlcHSupportIndOk() (*bool, bool)`

GetOlcHSupportIndOk returns a tuple with the OlcHSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOlcHSupportInd

`func (o *NfProfile) SetOlcHSupportInd(v bool)`

SetOlcHSupportInd sets OlcHSupportInd field to given value.

### HasOlcHSupportInd

`func (o *NfProfile) HasOlcHSupportInd() bool`

HasOlcHSupportInd returns a boolean if a field has been set.

### GetNfSetRecoveryTimeList

`func (o *NfProfile) GetNfSetRecoveryTimeList() map[string]time.Time`

GetNfSetRecoveryTimeList returns the NfSetRecoveryTimeList field if non-nil, zero value otherwise.

### GetNfSetRecoveryTimeListOk

`func (o *NfProfile) GetNfSetRecoveryTimeListOk() (*map[string]time.Time, bool)`

GetNfSetRecoveryTimeListOk returns a tuple with the NfSetRecoveryTimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetRecoveryTimeList

`func (o *NfProfile) SetNfSetRecoveryTimeList(v map[string]time.Time)`

SetNfSetRecoveryTimeList sets NfSetRecoveryTimeList field to given value.

### HasNfSetRecoveryTimeList

`func (o *NfProfile) HasNfSetRecoveryTimeList() bool`

HasNfSetRecoveryTimeList returns a boolean if a field has been set.

### GetServiceSetRecoveryTimeList

`func (o *NfProfile) GetServiceSetRecoveryTimeList() map[string]time.Time`

GetServiceSetRecoveryTimeList returns the ServiceSetRecoveryTimeList field if non-nil, zero value otherwise.

### GetServiceSetRecoveryTimeListOk

`func (o *NfProfile) GetServiceSetRecoveryTimeListOk() (*map[string]time.Time, bool)`

GetServiceSetRecoveryTimeListOk returns a tuple with the ServiceSetRecoveryTimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSetRecoveryTimeList

`func (o *NfProfile) SetServiceSetRecoveryTimeList(v map[string]time.Time)`

SetServiceSetRecoveryTimeList sets ServiceSetRecoveryTimeList field to given value.

### HasServiceSetRecoveryTimeList

`func (o *NfProfile) HasServiceSetRecoveryTimeList() bool`

HasServiceSetRecoveryTimeList returns a boolean if a field has been set.

### GetScpDomains

`func (o *NfProfile) GetScpDomains() []string`

GetScpDomains returns the ScpDomains field if non-nil, zero value otherwise.

### GetScpDomainsOk

`func (o *NfProfile) GetScpDomainsOk() (*[]string, bool)`

GetScpDomainsOk returns a tuple with the ScpDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpDomains

`func (o *NfProfile) SetScpDomains(v []string)`

SetScpDomains sets ScpDomains field to given value.

### HasScpDomains

`func (o *NfProfile) HasScpDomains() bool`

HasScpDomains returns a boolean if a field has been set.

### GetScpInfo

`func (o *NfProfile) GetScpInfo() ScpInfo`

GetScpInfo returns the ScpInfo field if non-nil, zero value otherwise.

### GetScpInfoOk

`func (o *NfProfile) GetScpInfoOk() (*ScpInfo, bool)`

GetScpInfoOk returns a tuple with the ScpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpInfo

`func (o *NfProfile) SetScpInfo(v ScpInfo)`

SetScpInfo sets ScpInfo field to given value.

### HasScpInfo

`func (o *NfProfile) HasScpInfo() bool`

HasScpInfo returns a boolean if a field has been set.

### GetSeppInfo

`func (o *NfProfile) GetSeppInfo() SeppInfo`

GetSeppInfo returns the SeppInfo field if non-nil, zero value otherwise.

### GetSeppInfoOk

`func (o *NfProfile) GetSeppInfoOk() (*SeppInfo, bool)`

GetSeppInfoOk returns a tuple with the SeppInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeppInfo

`func (o *NfProfile) SetSeppInfo(v SeppInfo)`

SetSeppInfo sets SeppInfo field to given value.

### HasSeppInfo

`func (o *NfProfile) HasSeppInfo() bool`

HasSeppInfo returns a boolean if a field has been set.

### GetVendorId

`func (o *NfProfile) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *NfProfile) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *NfProfile) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *NfProfile) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetSupportedVendorSpecificFeatures

`func (o *NfProfile) GetSupportedVendorSpecificFeatures() map[string][]VendorSpecificFeature`

GetSupportedVendorSpecificFeatures returns the SupportedVendorSpecificFeatures field if non-nil, zero value otherwise.

### GetSupportedVendorSpecificFeaturesOk

`func (o *NfProfile) GetSupportedVendorSpecificFeaturesOk() (*map[string][]VendorSpecificFeature, bool)`

GetSupportedVendorSpecificFeaturesOk returns a tuple with the SupportedVendorSpecificFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVendorSpecificFeatures

`func (o *NfProfile) SetSupportedVendorSpecificFeatures(v map[string][]VendorSpecificFeature)`

SetSupportedVendorSpecificFeatures sets SupportedVendorSpecificFeatures field to given value.

### HasSupportedVendorSpecificFeatures

`func (o *NfProfile) HasSupportedVendorSpecificFeatures() bool`

HasSupportedVendorSpecificFeatures returns a boolean if a field has been set.

### GetAanfInfoList

`func (o *NfProfile) GetAanfInfoList() map[string]AanfInfo`

GetAanfInfoList returns the AanfInfoList field if non-nil, zero value otherwise.

### GetAanfInfoListOk

`func (o *NfProfile) GetAanfInfoListOk() (*map[string]AanfInfo, bool)`

GetAanfInfoListOk returns a tuple with the AanfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAanfInfoList

`func (o *NfProfile) SetAanfInfoList(v map[string]AanfInfo)`

SetAanfInfoList sets AanfInfoList field to given value.

### HasAanfInfoList

`func (o *NfProfile) HasAanfInfoList() bool`

HasAanfInfoList returns a boolean if a field has been set.

### GetVar5gDdnmfInfo

`func (o *NfProfile) GetVar5gDdnmfInfo() Model5GDdnmfInfo`

GetVar5gDdnmfInfo returns the Var5gDdnmfInfo field if non-nil, zero value otherwise.

### GetVar5gDdnmfInfoOk

`func (o *NfProfile) GetVar5gDdnmfInfoOk() (*Model5GDdnmfInfo, bool)`

GetVar5gDdnmfInfoOk returns a tuple with the Var5gDdnmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gDdnmfInfo

`func (o *NfProfile) SetVar5gDdnmfInfo(v Model5GDdnmfInfo)`

SetVar5gDdnmfInfo sets Var5gDdnmfInfo field to given value.

### HasVar5gDdnmfInfo

`func (o *NfProfile) HasVar5gDdnmfInfo() bool`

HasVar5gDdnmfInfo returns a boolean if a field has been set.

### GetMfafInfo

`func (o *NfProfile) GetMfafInfo() MfafInfo`

GetMfafInfo returns the MfafInfo field if non-nil, zero value otherwise.

### GetMfafInfoOk

`func (o *NfProfile) GetMfafInfoOk() (*MfafInfo, bool)`

GetMfafInfoOk returns a tuple with the MfafInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfafInfo

`func (o *NfProfile) SetMfafInfo(v MfafInfo)`

SetMfafInfo sets MfafInfo field to given value.

### HasMfafInfo

`func (o *NfProfile) HasMfafInfo() bool`

HasMfafInfo returns a boolean if a field has been set.

### GetEasdfInfoList

`func (o *NfProfile) GetEasdfInfoList() map[string]EasdfInfo`

GetEasdfInfoList returns the EasdfInfoList field if non-nil, zero value otherwise.

### GetEasdfInfoListOk

`func (o *NfProfile) GetEasdfInfoListOk() (*map[string]EasdfInfo, bool)`

GetEasdfInfoListOk returns a tuple with the EasdfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasdfInfoList

`func (o *NfProfile) SetEasdfInfoList(v map[string]EasdfInfo)`

SetEasdfInfoList sets EasdfInfoList field to given value.

### HasEasdfInfoList

`func (o *NfProfile) HasEasdfInfoList() bool`

HasEasdfInfoList returns a boolean if a field has been set.

### GetDccfInfo

`func (o *NfProfile) GetDccfInfo() DccfInfo`

GetDccfInfo returns the DccfInfo field if non-nil, zero value otherwise.

### GetDccfInfoOk

`func (o *NfProfile) GetDccfInfoOk() (*DccfInfo, bool)`

GetDccfInfoOk returns a tuple with the DccfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccfInfo

`func (o *NfProfile) SetDccfInfo(v DccfInfo)`

SetDccfInfo sets DccfInfo field to given value.

### HasDccfInfo

`func (o *NfProfile) HasDccfInfo() bool`

HasDccfInfo returns a boolean if a field has been set.

### GetNsacfInfoList

`func (o *NfProfile) GetNsacfInfoList() map[string]NsacfInfo`

GetNsacfInfoList returns the NsacfInfoList field if non-nil, zero value otherwise.

### GetNsacfInfoListOk

`func (o *NfProfile) GetNsacfInfoListOk() (*map[string]NsacfInfo, bool)`

GetNsacfInfoListOk returns a tuple with the NsacfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsacfInfoList

`func (o *NfProfile) SetNsacfInfoList(v map[string]NsacfInfo)`

SetNsacfInfoList sets NsacfInfoList field to given value.

### HasNsacfInfoList

`func (o *NfProfile) HasNsacfInfoList() bool`

HasNsacfInfoList returns a boolean if a field has been set.

### GetMbSmfInfoList

`func (o *NfProfile) GetMbSmfInfoList() map[string]MbSmfInfo`

GetMbSmfInfoList returns the MbSmfInfoList field if non-nil, zero value otherwise.

### GetMbSmfInfoListOk

`func (o *NfProfile) GetMbSmfInfoListOk() (*map[string]MbSmfInfo, bool)`

GetMbSmfInfoListOk returns a tuple with the MbSmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbSmfInfoList

`func (o *NfProfile) SetMbSmfInfoList(v map[string]MbSmfInfo)`

SetMbSmfInfoList sets MbSmfInfoList field to given value.

### HasMbSmfInfoList

`func (o *NfProfile) HasMbSmfInfoList() bool`

HasMbSmfInfoList returns a boolean if a field has been set.

### GetTsctsfInfoList

`func (o *NfProfile) GetTsctsfInfoList() map[string]TsctsfInfo`

GetTsctsfInfoList returns the TsctsfInfoList field if non-nil, zero value otherwise.

### GetTsctsfInfoListOk

`func (o *NfProfile) GetTsctsfInfoListOk() (*map[string]TsctsfInfo, bool)`

GetTsctsfInfoListOk returns a tuple with the TsctsfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsctsfInfoList

`func (o *NfProfile) SetTsctsfInfoList(v map[string]TsctsfInfo)`

SetTsctsfInfoList sets TsctsfInfoList field to given value.

### HasTsctsfInfoList

`func (o *NfProfile) HasTsctsfInfoList() bool`

HasTsctsfInfoList returns a boolean if a field has been set.

### GetMbUpfInfoList

`func (o *NfProfile) GetMbUpfInfoList() map[string]MbUpfInfo`

GetMbUpfInfoList returns the MbUpfInfoList field if non-nil, zero value otherwise.

### GetMbUpfInfoListOk

`func (o *NfProfile) GetMbUpfInfoListOk() (*map[string]MbUpfInfo, bool)`

GetMbUpfInfoListOk returns a tuple with the MbUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbUpfInfoList

`func (o *NfProfile) SetMbUpfInfoList(v map[string]MbUpfInfo)`

SetMbUpfInfoList sets MbUpfInfoList field to given value.

### HasMbUpfInfoList

`func (o *NfProfile) HasMbUpfInfoList() bool`

HasMbUpfInfoList returns a boolean if a field has been set.

### GetTrustAfInfo

`func (o *NfProfile) GetTrustAfInfo() TrustAfInfo`

GetTrustAfInfo returns the TrustAfInfo field if non-nil, zero value otherwise.

### GetTrustAfInfoOk

`func (o *NfProfile) GetTrustAfInfoOk() (*TrustAfInfo, bool)`

GetTrustAfInfoOk returns a tuple with the TrustAfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAfInfo

`func (o *NfProfile) SetTrustAfInfo(v TrustAfInfo)`

SetTrustAfInfo sets TrustAfInfo field to given value.

### HasTrustAfInfo

`func (o *NfProfile) HasTrustAfInfo() bool`

HasTrustAfInfo returns a boolean if a field has been set.

### GetNssaafInfo

`func (o *NfProfile) GetNssaafInfo() NssaafInfo`

GetNssaafInfo returns the NssaafInfo field if non-nil, zero value otherwise.

### GetNssaafInfoOk

`func (o *NfProfile) GetNssaafInfoOk() (*NssaafInfo, bool)`

GetNssaafInfoOk returns a tuple with the NssaafInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssaafInfo

`func (o *NfProfile) SetNssaafInfo(v NssaafInfo)`

SetNssaafInfo sets NssaafInfo field to given value.

### HasNssaafInfo

`func (o *NfProfile) HasNssaafInfo() bool`

HasNssaafInfo returns a boolean if a field has been set.

### GetHniList

`func (o *NfProfile) GetHniList() []string`

GetHniList returns the HniList field if non-nil, zero value otherwise.

### GetHniListOk

`func (o *NfProfile) GetHniListOk() (*[]string, bool)`

GetHniListOk returns a tuple with the HniList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHniList

`func (o *NfProfile) SetHniList(v []string)`

SetHniList sets HniList field to given value.

### HasHniList

`func (o *NfProfile) HasHniList() bool`

HasHniList returns a boolean if a field has been set.

### GetIwmscInfo

`func (o *NfProfile) GetIwmscInfo() IwmscInfo`

GetIwmscInfo returns the IwmscInfo field if non-nil, zero value otherwise.

### GetIwmscInfoOk

`func (o *NfProfile) GetIwmscInfoOk() (*IwmscInfo, bool)`

GetIwmscInfoOk returns a tuple with the IwmscInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwmscInfo

`func (o *NfProfile) SetIwmscInfo(v IwmscInfo)`

SetIwmscInfo sets IwmscInfo field to given value.

### HasIwmscInfo

`func (o *NfProfile) HasIwmscInfo() bool`

HasIwmscInfo returns a boolean if a field has been set.

### GetMnpfInfo

`func (o *NfProfile) GetMnpfInfo() MnpfInfo`

GetMnpfInfo returns the MnpfInfo field if non-nil, zero value otherwise.

### GetMnpfInfoOk

`func (o *NfProfile) GetMnpfInfoOk() (*MnpfInfo, bool)`

GetMnpfInfoOk returns a tuple with the MnpfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnpfInfo

`func (o *NfProfile) SetMnpfInfo(v MnpfInfo)`

SetMnpfInfo sets MnpfInfo field to given value.

### HasMnpfInfo

`func (o *NfProfile) HasMnpfInfo() bool`

HasMnpfInfo returns a boolean if a field has been set.

### GetSmsfInfo

`func (o *NfProfile) GetSmsfInfo() SmsfInfo`

GetSmsfInfo returns the SmsfInfo field if non-nil, zero value otherwise.

### GetSmsfInfoOk

`func (o *NfProfile) GetSmsfInfoOk() (*SmsfInfo, bool)`

GetSmsfInfoOk returns a tuple with the SmsfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfInfo

`func (o *NfProfile) SetSmsfInfo(v SmsfInfo)`

SetSmsfInfo sets SmsfInfo field to given value.

### HasSmsfInfo

`func (o *NfProfile) HasSmsfInfo() bool`

HasSmsfInfo returns a boolean if a field has been set.

### GetDcsfInfoList

`func (o *NfProfile) GetDcsfInfoList() map[string]DcsfInfo`

GetDcsfInfoList returns the DcsfInfoList field if non-nil, zero value otherwise.

### GetDcsfInfoListOk

`func (o *NfProfile) GetDcsfInfoListOk() (*map[string]DcsfInfo, bool)`

GetDcsfInfoListOk returns a tuple with the DcsfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcsfInfoList

`func (o *NfProfile) SetDcsfInfoList(v map[string]DcsfInfo)`

SetDcsfInfoList sets DcsfInfoList field to given value.

### HasDcsfInfoList

`func (o *NfProfile) HasDcsfInfoList() bool`

HasDcsfInfoList returns a boolean if a field has been set.

### GetMrfInfoList

`func (o *NfProfile) GetMrfInfoList() map[string]MrfInfo`

GetMrfInfoList returns the MrfInfoList field if non-nil, zero value otherwise.

### GetMrfInfoListOk

`func (o *NfProfile) GetMrfInfoListOk() (*map[string]MrfInfo, bool)`

GetMrfInfoListOk returns a tuple with the MrfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrfInfoList

`func (o *NfProfile) SetMrfInfoList(v map[string]MrfInfo)`

SetMrfInfoList sets MrfInfoList field to given value.

### HasMrfInfoList

`func (o *NfProfile) HasMrfInfoList() bool`

HasMrfInfoList returns a boolean if a field has been set.

### GetMrfpInfoList

`func (o *NfProfile) GetMrfpInfoList() map[string]MrfpInfo`

GetMrfpInfoList returns the MrfpInfoList field if non-nil, zero value otherwise.

### GetMrfpInfoListOk

`func (o *NfProfile) GetMrfpInfoListOk() (*map[string]MrfpInfo, bool)`

GetMrfpInfoListOk returns a tuple with the MrfpInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrfpInfoList

`func (o *NfProfile) SetMrfpInfoList(v map[string]MrfpInfo)`

SetMrfpInfoList sets MrfpInfoList field to given value.

### HasMrfpInfoList

`func (o *NfProfile) HasMrfpInfoList() bool`

HasMrfpInfoList returns a boolean if a field has been set.

### GetMfInfoList

`func (o *NfProfile) GetMfInfoList() map[string]MfInfo`

GetMfInfoList returns the MfInfoList field if non-nil, zero value otherwise.

### GetMfInfoListOk

`func (o *NfProfile) GetMfInfoListOk() (*map[string]MfInfo, bool)`

GetMfInfoListOk returns a tuple with the MfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfInfoList

`func (o *NfProfile) SetMfInfoList(v map[string]MfInfo)`

SetMfInfoList sets MfInfoList field to given value.

### HasMfInfoList

`func (o *NfProfile) HasMfInfoList() bool`

HasMfInfoList returns a boolean if a field has been set.

### GetAdrfInfoList

`func (o *NfProfile) GetAdrfInfoList() map[string]AdrfInfo`

GetAdrfInfoList returns the AdrfInfoList field if non-nil, zero value otherwise.

### GetAdrfInfoListOk

`func (o *NfProfile) GetAdrfInfoListOk() (*map[string]AdrfInfo, bool)`

GetAdrfInfoListOk returns a tuple with the AdrfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfInfoList

`func (o *NfProfile) SetAdrfInfoList(v map[string]AdrfInfo)`

SetAdrfInfoList sets AdrfInfoList field to given value.

### HasAdrfInfoList

`func (o *NfProfile) HasAdrfInfoList() bool`

HasAdrfInfoList returns a boolean if a field has been set.

### GetSelectionConditions

`func (o *NfProfile) GetSelectionConditions() SelectionConditions`

GetSelectionConditions returns the SelectionConditions field if non-nil, zero value otherwise.

### GetSelectionConditionsOk

`func (o *NfProfile) GetSelectionConditionsOk() (*SelectionConditions, bool)`

GetSelectionConditionsOk returns a tuple with the SelectionConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionConditions

`func (o *NfProfile) SetSelectionConditions(v SelectionConditions)`

SetSelectionConditions sets SelectionConditions field to given value.

### HasSelectionConditions

`func (o *NfProfile) HasSelectionConditions() bool`

HasSelectionConditions returns a boolean if a field has been set.

### GetCanaryRelease

`func (o *NfProfile) GetCanaryRelease() bool`

GetCanaryRelease returns the CanaryRelease field if non-nil, zero value otherwise.

### GetCanaryReleaseOk

`func (o *NfProfile) GetCanaryReleaseOk() (*bool, bool)`

GetCanaryReleaseOk returns a tuple with the CanaryRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryRelease

`func (o *NfProfile) SetCanaryRelease(v bool)`

SetCanaryRelease sets CanaryRelease field to given value.

### HasCanaryRelease

`func (o *NfProfile) HasCanaryRelease() bool`

HasCanaryRelease returns a boolean if a field has been set.

### GetExclusiveCanaryReleaseSelection

`func (o *NfProfile) GetExclusiveCanaryReleaseSelection() bool`

GetExclusiveCanaryReleaseSelection returns the ExclusiveCanaryReleaseSelection field if non-nil, zero value otherwise.

### GetExclusiveCanaryReleaseSelectionOk

`func (o *NfProfile) GetExclusiveCanaryReleaseSelectionOk() (*bool, bool)`

GetExclusiveCanaryReleaseSelectionOk returns a tuple with the ExclusiveCanaryReleaseSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveCanaryReleaseSelection

`func (o *NfProfile) SetExclusiveCanaryReleaseSelection(v bool)`

SetExclusiveCanaryReleaseSelection sets ExclusiveCanaryReleaseSelection field to given value.

### HasExclusiveCanaryReleaseSelection

`func (o *NfProfile) HasExclusiveCanaryReleaseSelection() bool`

HasExclusiveCanaryReleaseSelection returns a boolean if a field has been set.

### GetSharedProfileDataId

`func (o *NfProfile) GetSharedProfileDataId() string`

GetSharedProfileDataId returns the SharedProfileDataId field if non-nil, zero value otherwise.

### GetSharedProfileDataIdOk

`func (o *NfProfile) GetSharedProfileDataIdOk() (*string, bool)`

GetSharedProfileDataIdOk returns a tuple with the SharedProfileDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedProfileDataId

`func (o *NfProfile) SetSharedProfileDataId(v string)`

SetSharedProfileDataId sets SharedProfileDataId field to given value.

### HasSharedProfileDataId

`func (o *NfProfile) HasSharedProfileDataId() bool`

HasSharedProfileDataId returns a boolean if a field has been set.

### GetShutdownTime

`func (o *NfProfile) GetShutdownTime() time.Time`

GetShutdownTime returns the ShutdownTime field if non-nil, zero value otherwise.

### GetShutdownTimeOk

`func (o *NfProfile) GetShutdownTimeOk() (*time.Time, bool)`

GetShutdownTimeOk returns a tuple with the ShutdownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownTime

`func (o *NfProfile) SetShutdownTime(v time.Time)`

SetShutdownTime sets ShutdownTime field to given value.

### HasShutdownTime

`func (o *NfProfile) HasShutdownTime() bool`

HasShutdownTime returns a boolean if a field has been set.

### GetSupportedRcfs

`func (o *NfProfile) GetSupportedRcfs() []string`

GetSupportedRcfs returns the SupportedRcfs field if non-nil, zero value otherwise.

### GetSupportedRcfsOk

`func (o *NfProfile) GetSupportedRcfsOk() (*[]string, bool)`

GetSupportedRcfsOk returns a tuple with the SupportedRcfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRcfs

`func (o *NfProfile) SetSupportedRcfs(v []string)`

SetSupportedRcfs sets SupportedRcfs field to given value.

### HasSupportedRcfs

`func (o *NfProfile) HasSupportedRcfs() bool`

HasSupportedRcfs returns a boolean if a field has been set.

### GetCanaryPrecedenceOverPreferred

`func (o *NfProfile) GetCanaryPrecedenceOverPreferred() bool`

GetCanaryPrecedenceOverPreferred returns the CanaryPrecedenceOverPreferred field if non-nil, zero value otherwise.

### GetCanaryPrecedenceOverPreferredOk

`func (o *NfProfile) GetCanaryPrecedenceOverPreferredOk() (*bool, bool)`

GetCanaryPrecedenceOverPreferredOk returns a tuple with the CanaryPrecedenceOverPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryPrecedenceOverPreferred

`func (o *NfProfile) SetCanaryPrecedenceOverPreferred(v bool)`

SetCanaryPrecedenceOverPreferred sets CanaryPrecedenceOverPreferred field to given value.

### HasCanaryPrecedenceOverPreferred

`func (o *NfProfile) HasCanaryPrecedenceOverPreferred() bool`

HasCanaryPrecedenceOverPreferred returns a boolean if a field has been set.

### GetImsasInfoList

`func (o *NfProfile) GetImsasInfoList() map[string]ImsasInfo`

GetImsasInfoList returns the ImsasInfoList field if non-nil, zero value otherwise.

### GetImsasInfoListOk

`func (o *NfProfile) GetImsasInfoListOk() (*map[string]ImsasInfo, bool)`

GetImsasInfoListOk returns a tuple with the ImsasInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsasInfoList

`func (o *NfProfile) SetImsasInfoList(v map[string]ImsasInfo)`

SetImsasInfoList sets ImsasInfoList field to given value.

### HasImsasInfoList

`func (o *NfProfile) HasImsasInfoList() bool`

HasImsasInfoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


