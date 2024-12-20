// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos.Outputs
{

    [OutputType]
    public sealed class PacketFilterConfigPacketFilterConfig
    {
        public readonly string Action;
        public readonly int? Depth;
        public readonly int? Depth2;
        public readonly int DportEnd;
        public readonly int DportStart;
        public readonly int? IsNot;
        public readonly int? IsNot2;
        public readonly string? MatchBegin;
        public readonly string? MatchBegin2;
        public readonly string? MatchLogic;
        public readonly string? MatchType;
        public readonly string? MatchType2;
        public readonly int? Offset;
        public readonly int? Offset2;
        public readonly int? PktLenGt;
        public readonly int PktlenMax;
        public readonly int PktlenMin;
        public readonly string Protocol;
        public readonly int SportEnd;
        public readonly int SportStart;
        public readonly string? Str;
        public readonly string? Str2;

        [OutputConstructor]
        private PacketFilterConfigPacketFilterConfig(
            string action,

            int? depth,

            int? depth2,

            int dportEnd,

            int dportStart,

            int? isNot,

            int? isNot2,

            string? matchBegin,

            string? matchBegin2,

            string? matchLogic,

            string? matchType,

            string? matchType2,

            int? offset,

            int? offset2,

            int? pktLenGt,

            int pktlenMax,

            int pktlenMin,

            string protocol,

            int sportEnd,

            int sportStart,

            string? str,

            string? str2)
        {
            Action = action;
            Depth = depth;
            Depth2 = depth2;
            DportEnd = dportEnd;
            DportStart = dportStart;
            IsNot = isNot;
            IsNot2 = isNot2;
            MatchBegin = matchBegin;
            MatchBegin2 = matchBegin2;
            MatchLogic = matchLogic;
            MatchType = matchType;
            MatchType2 = matchType2;
            Offset = offset;
            Offset2 = offset2;
            PktLenGt = pktLenGt;
            PktlenMax = pktlenMax;
            PktlenMin = pktlenMin;
            Protocol = protocol;
            SportEnd = sportEnd;
            SportStart = sportStart;
            Str = str;
            Str2 = str2;
        }
    }
}
