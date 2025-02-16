# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from ._inputs import *

__all__ = ['Provider']


class Provider(pulumi.ProviderResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_material: Optional[pulumi.Input[str]] = None,
                 cert_material: Optional[pulumi.Input[str]] = None,
                 cert_path: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 key_material: Optional[pulumi.Input[str]] = None,
                 registry_auth: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderRegistryAuthArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the docker package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_material: PEM-encoded content of Docker host CA certificate
        :param pulumi.Input[str] cert_material: PEM-encoded content of Docker client certificate
        :param pulumi.Input[str] cert_path: Path to directory with Docker TLS config
        :param pulumi.Input[str] host: The Docker daemon address
        :param pulumi.Input[str] key_material: PEM-encoded content of Docker client private key
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['ca_material'] = ca_material
            __props__['cert_material'] = cert_material
            __props__['cert_path'] = cert_path
            if host is None:
                host = (_utilities.get_env('DOCKER_HOST') or 'unix:///var/run/docker.sock')
            __props__['host'] = host
            __props__['key_material'] = key_material
            __props__['registry_auth'] = pulumi.Output.from_input(registry_auth).apply(pulumi.runtime.to_json) if registry_auth is not None else None
        super(Provider, __self__).__init__(
            'docker',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

