# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetCrossBorderComplianceResult',
    'AwaitableGetCrossBorderComplianceResult',
    'get_cross_border_compliance',
    'get_cross_border_compliance_output',
]

@pulumi.output_type
class GetCrossBorderComplianceResult:
    """
    A collection of values returned by getCrossBorderCompliance.
    """
    def __init__(__self__, business_address=None, company=None, compliance_id=None, email=None, id=None, issuing_authority=None, legal_person=None, manager=None, manager_address=None, manager_id=None, manager_telephone=None, post_code=None, result_output_file=None, service_end_date=None, service_provider=None, service_start_date=None, state=None, uniform_social_credit_code=None):
        if business_address and not isinstance(business_address, str):
            raise TypeError("Expected argument 'business_address' to be a str")
        pulumi.set(__self__, "business_address", business_address)
        if company and not isinstance(company, str):
            raise TypeError("Expected argument 'company' to be a str")
        pulumi.set(__self__, "company", company)
        if compliance_id and not isinstance(compliance_id, int):
            raise TypeError("Expected argument 'compliance_id' to be a int")
        pulumi.set(__self__, "compliance_id", compliance_id)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if issuing_authority and not isinstance(issuing_authority, str):
            raise TypeError("Expected argument 'issuing_authority' to be a str")
        pulumi.set(__self__, "issuing_authority", issuing_authority)
        if legal_person and not isinstance(legal_person, str):
            raise TypeError("Expected argument 'legal_person' to be a str")
        pulumi.set(__self__, "legal_person", legal_person)
        if manager and not isinstance(manager, str):
            raise TypeError("Expected argument 'manager' to be a str")
        pulumi.set(__self__, "manager", manager)
        if manager_address and not isinstance(manager_address, str):
            raise TypeError("Expected argument 'manager_address' to be a str")
        pulumi.set(__self__, "manager_address", manager_address)
        if manager_id and not isinstance(manager_id, str):
            raise TypeError("Expected argument 'manager_id' to be a str")
        pulumi.set(__self__, "manager_id", manager_id)
        if manager_telephone and not isinstance(manager_telephone, str):
            raise TypeError("Expected argument 'manager_telephone' to be a str")
        pulumi.set(__self__, "manager_telephone", manager_telephone)
        if post_code and not isinstance(post_code, int):
            raise TypeError("Expected argument 'post_code' to be a int")
        pulumi.set(__self__, "post_code", post_code)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if service_end_date and not isinstance(service_end_date, str):
            raise TypeError("Expected argument 'service_end_date' to be a str")
        pulumi.set(__self__, "service_end_date", service_end_date)
        if service_provider and not isinstance(service_provider, str):
            raise TypeError("Expected argument 'service_provider' to be a str")
        pulumi.set(__self__, "service_provider", service_provider)
        if service_start_date and not isinstance(service_start_date, str):
            raise TypeError("Expected argument 'service_start_date' to be a str")
        pulumi.set(__self__, "service_start_date", service_start_date)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if uniform_social_credit_code and not isinstance(uniform_social_credit_code, str):
            raise TypeError("Expected argument 'uniform_social_credit_code' to be a str")
        pulumi.set(__self__, "uniform_social_credit_code", uniform_social_credit_code)

    @property
    @pulumi.getter(name="businessAddress")
    def business_address(self) -> Optional[str]:
        return pulumi.get(self, "business_address")

    @property
    @pulumi.getter
    def company(self) -> Optional[str]:
        return pulumi.get(self, "company")

    @property
    @pulumi.getter(name="complianceId")
    def compliance_id(self) -> Optional[int]:
        return pulumi.get(self, "compliance_id")

    @property
    @pulumi.getter
    def email(self) -> Optional[str]:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="issuingAuthority")
    def issuing_authority(self) -> Optional[str]:
        return pulumi.get(self, "issuing_authority")

    @property
    @pulumi.getter(name="legalPerson")
    def legal_person(self) -> Optional[str]:
        return pulumi.get(self, "legal_person")

    @property
    @pulumi.getter
    def manager(self) -> Optional[str]:
        return pulumi.get(self, "manager")

    @property
    @pulumi.getter(name="managerAddress")
    def manager_address(self) -> Optional[str]:
        return pulumi.get(self, "manager_address")

    @property
    @pulumi.getter(name="managerId")
    def manager_id(self) -> Optional[str]:
        return pulumi.get(self, "manager_id")

    @property
    @pulumi.getter(name="managerTelephone")
    def manager_telephone(self) -> Optional[str]:
        return pulumi.get(self, "manager_telephone")

    @property
    @pulumi.getter(name="postCode")
    def post_code(self) -> Optional[int]:
        return pulumi.get(self, "post_code")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="serviceEndDate")
    def service_end_date(self) -> Optional[str]:
        return pulumi.get(self, "service_end_date")

    @property
    @pulumi.getter(name="serviceProvider")
    def service_provider(self) -> Optional[str]:
        return pulumi.get(self, "service_provider")

    @property
    @pulumi.getter(name="serviceStartDate")
    def service_start_date(self) -> Optional[str]:
        return pulumi.get(self, "service_start_date")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="uniformSocialCreditCode")
    def uniform_social_credit_code(self) -> Optional[str]:
        return pulumi.get(self, "uniform_social_credit_code")


class AwaitableGetCrossBorderComplianceResult(GetCrossBorderComplianceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCrossBorderComplianceResult(
            business_address=self.business_address,
            company=self.company,
            compliance_id=self.compliance_id,
            email=self.email,
            id=self.id,
            issuing_authority=self.issuing_authority,
            legal_person=self.legal_person,
            manager=self.manager,
            manager_address=self.manager_address,
            manager_id=self.manager_id,
            manager_telephone=self.manager_telephone,
            post_code=self.post_code,
            result_output_file=self.result_output_file,
            service_end_date=self.service_end_date,
            service_provider=self.service_provider,
            service_start_date=self.service_start_date,
            state=self.state,
            uniform_social_credit_code=self.uniform_social_credit_code)


def get_cross_border_compliance(business_address: Optional[str] = None,
                                company: Optional[str] = None,
                                compliance_id: Optional[int] = None,
                                email: Optional[str] = None,
                                issuing_authority: Optional[str] = None,
                                legal_person: Optional[str] = None,
                                manager: Optional[str] = None,
                                manager_address: Optional[str] = None,
                                manager_id: Optional[str] = None,
                                manager_telephone: Optional[str] = None,
                                post_code: Optional[int] = None,
                                result_output_file: Optional[str] = None,
                                service_end_date: Optional[str] = None,
                                service_provider: Optional[str] = None,
                                service_start_date: Optional[str] = None,
                                state: Optional[str] = None,
                                uniform_social_credit_code: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCrossBorderComplianceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['businessAddress'] = business_address
    __args__['company'] = company
    __args__['complianceId'] = compliance_id
    __args__['email'] = email
    __args__['issuingAuthority'] = issuing_authority
    __args__['legalPerson'] = legal_person
    __args__['manager'] = manager
    __args__['managerAddress'] = manager_address
    __args__['managerId'] = manager_id
    __args__['managerTelephone'] = manager_telephone
    __args__['postCode'] = post_code
    __args__['resultOutputFile'] = result_output_file
    __args__['serviceEndDate'] = service_end_date
    __args__['serviceProvider'] = service_provider
    __args__['serviceStartDate'] = service_start_date
    __args__['state'] = state
    __args__['uniformSocialCreditCode'] = uniform_social_credit_code
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ccn/getCrossBorderCompliance:getCrossBorderCompliance', __args__, opts=opts, typ=GetCrossBorderComplianceResult).value

    return AwaitableGetCrossBorderComplianceResult(
        business_address=pulumi.get(__ret__, 'business_address'),
        company=pulumi.get(__ret__, 'company'),
        compliance_id=pulumi.get(__ret__, 'compliance_id'),
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        issuing_authority=pulumi.get(__ret__, 'issuing_authority'),
        legal_person=pulumi.get(__ret__, 'legal_person'),
        manager=pulumi.get(__ret__, 'manager'),
        manager_address=pulumi.get(__ret__, 'manager_address'),
        manager_id=pulumi.get(__ret__, 'manager_id'),
        manager_telephone=pulumi.get(__ret__, 'manager_telephone'),
        post_code=pulumi.get(__ret__, 'post_code'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        service_end_date=pulumi.get(__ret__, 'service_end_date'),
        service_provider=pulumi.get(__ret__, 'service_provider'),
        service_start_date=pulumi.get(__ret__, 'service_start_date'),
        state=pulumi.get(__ret__, 'state'),
        uniform_social_credit_code=pulumi.get(__ret__, 'uniform_social_credit_code'))


@_utilities.lift_output_func(get_cross_border_compliance)
def get_cross_border_compliance_output(business_address: Optional[pulumi.Input[Optional[str]]] = None,
                                       company: Optional[pulumi.Input[Optional[str]]] = None,
                                       compliance_id: Optional[pulumi.Input[Optional[int]]] = None,
                                       email: Optional[pulumi.Input[Optional[str]]] = None,
                                       issuing_authority: Optional[pulumi.Input[Optional[str]]] = None,
                                       legal_person: Optional[pulumi.Input[Optional[str]]] = None,
                                       manager: Optional[pulumi.Input[Optional[str]]] = None,
                                       manager_address: Optional[pulumi.Input[Optional[str]]] = None,
                                       manager_id: Optional[pulumi.Input[Optional[str]]] = None,
                                       manager_telephone: Optional[pulumi.Input[Optional[str]]] = None,
                                       post_code: Optional[pulumi.Input[Optional[int]]] = None,
                                       result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                       service_end_date: Optional[pulumi.Input[Optional[str]]] = None,
                                       service_provider: Optional[pulumi.Input[Optional[str]]] = None,
                                       service_start_date: Optional[pulumi.Input[Optional[str]]] = None,
                                       state: Optional[pulumi.Input[Optional[str]]] = None,
                                       uniform_social_credit_code: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCrossBorderComplianceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...