// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcmq
{
    [TencentcloudResourceType("tencentcloud:Tcmq/subscribe:Subscribe")]
    public partial class Subscribe : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field
        /// indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots
        /// (namely up to 16 segments).
        /// </summary>
        [Output("bindingKeys")]
        public Output<ImmutableArray<string>> BindingKeys { get; private set; } = null!;

        /// <summary>
        /// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with
        /// `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service
        /// cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages
        /// will not be received. Currently, messages can be pushed only to the public network and classic network.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16
        /// characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag`
        /// is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the
        /// topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists
        /// in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive
        /// messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then
        /// no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag`
        /// and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the
        /// subscriber.
        /// </summary>
        [Output("filterTags")]
        public Output<ImmutableArray<string>> FilterTags { get; private set; } = null!;

        /// <summary>
        /// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value
        /// must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
        /// </summary>
        [Output("notifyContentFormat")]
        public Output<string?> NotifyContentFormat { get; private set; } = null!;

        /// <summary>
        /// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1.
        /// `BACKOFF_RETRY`: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of
        /// retries, and continue to push the next message; 2. `EXPONENTIAL_DECAY_RETRY`: exponential decay retry, which is to retry
        /// at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for
        /// one day, failed messages will be discarded at most after one day of retry. Default value: `EXPONENTIAL_DECAY_RETRY`.
        /// </summary>
        [Output("notifyStrategy")]
        public Output<string?> NotifyStrategy { get; private set; } = null!;

        /// <summary>
        /// ubscription protocol. Currently, two protocols are supported: `http` and `queue`. To use the `http` protocol, you need
        /// to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ
        /// queue and you can pull them concurrently.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up
        /// to 64 letters, digits, and hyphens and must begin with a letter.
        /// </summary>
        [Output("subscriptionName")]
        public Output<string> SubscriptionName { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64
        /// letters, digits, and hyphens and must begin with a letter.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;


        /// <summary>
        /// Create a Subscribe resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subscribe(string name, SubscribeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcmq/subscribe:Subscribe", name, args ?? new SubscribeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subscribe(string name, Input<string> id, SubscribeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcmq/subscribe:Subscribe", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/matrixorigin",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Subscribe resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subscribe Get(string name, Input<string> id, SubscribeState? state = null, CustomResourceOptions? options = null)
        {
            return new Subscribe(name, id, state, options);
        }
    }

    public sealed class SubscribeArgs : global::Pulumi.ResourceArgs
    {
        [Input("bindingKeys")]
        private InputList<string>? _bindingKeys;

        /// <summary>
        /// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field
        /// indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots
        /// (namely up to 16 segments).
        /// </summary>
        public InputList<string> BindingKeys
        {
            get => _bindingKeys ?? (_bindingKeys = new InputList<string>());
            set => _bindingKeys = value;
        }

        /// <summary>
        /// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with
        /// `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service
        /// cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages
        /// will not be received. Currently, messages can be pushed only to the public network and classic network.
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        [Input("filterTags")]
        private InputList<string>? _filterTags;

        /// <summary>
        /// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16
        /// characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag`
        /// is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the
        /// topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists
        /// in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive
        /// messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then
        /// no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag`
        /// and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the
        /// subscriber.
        /// </summary>
        public InputList<string> FilterTags
        {
            get => _filterTags ?? (_filterTags = new InputList<string>());
            set => _filterTags = value;
        }

        /// <summary>
        /// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value
        /// must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
        /// </summary>
        [Input("notifyContentFormat")]
        public Input<string>? NotifyContentFormat { get; set; }

        /// <summary>
        /// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1.
        /// `BACKOFF_RETRY`: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of
        /// retries, and continue to push the next message; 2. `EXPONENTIAL_DECAY_RETRY`: exponential decay retry, which is to retry
        /// at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for
        /// one day, failed messages will be discarded at most after one day of retry. Default value: `EXPONENTIAL_DECAY_RETRY`.
        /// </summary>
        [Input("notifyStrategy")]
        public Input<string>? NotifyStrategy { get; set; }

        /// <summary>
        /// ubscription protocol. Currently, two protocols are supported: `http` and `queue`. To use the `http` protocol, you need
        /// to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ
        /// queue and you can pull them concurrently.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up
        /// to 64 letters, digits, and hyphens and must begin with a letter.
        /// </summary>
        [Input("subscriptionName", required: true)]
        public Input<string> SubscriptionName { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64
        /// letters, digits, and hyphens and must begin with a letter.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public SubscribeArgs()
        {
        }
        public static new SubscribeArgs Empty => new SubscribeArgs();
    }

    public sealed class SubscribeState : global::Pulumi.ResourceArgs
    {
        [Input("bindingKeys")]
        private InputList<string>? _bindingKeys;

        /// <summary>
        /// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field
        /// indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots
        /// (namely up to 16 segments).
        /// </summary>
        public InputList<string> BindingKeys
        {
            get => _bindingKeys ?? (_bindingKeys = new InputList<string>());
            set => _bindingKeys = value;
        }

        /// <summary>
        /// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with
        /// `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service
        /// cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages
        /// will not be received. Currently, messages can be pushed only to the public network and classic network.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        [Input("filterTags")]
        private InputList<string>? _filterTags;

        /// <summary>
        /// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16
        /// characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag`
        /// is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the
        /// topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists
        /// in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive
        /// messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then
        /// no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag`
        /// and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the
        /// subscriber.
        /// </summary>
        public InputList<string> FilterTags
        {
            get => _filterTags ?? (_filterTags = new InputList<string>());
            set => _filterTags = value;
        }

        /// <summary>
        /// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value
        /// must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
        /// </summary>
        [Input("notifyContentFormat")]
        public Input<string>? NotifyContentFormat { get; set; }

        /// <summary>
        /// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1.
        /// `BACKOFF_RETRY`: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of
        /// retries, and continue to push the next message; 2. `EXPONENTIAL_DECAY_RETRY`: exponential decay retry, which is to retry
        /// at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for
        /// one day, failed messages will be discarded at most after one day of retry. Default value: `EXPONENTIAL_DECAY_RETRY`.
        /// </summary>
        [Input("notifyStrategy")]
        public Input<string>? NotifyStrategy { get; set; }

        /// <summary>
        /// ubscription protocol. Currently, two protocols are supported: `http` and `queue`. To use the `http` protocol, you need
        /// to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ
        /// queue and you can pull them concurrently.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up
        /// to 64 letters, digits, and hyphens and must begin with a letter.
        /// </summary>
        [Input("subscriptionName")]
        public Input<string>? SubscriptionName { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64
        /// letters, digits, and hyphens and must begin with a letter.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public SubscribeState()
        {
        }
        public static new SubscribeState Empty => new SubscribeState();
    }
}
