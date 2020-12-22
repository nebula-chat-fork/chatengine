// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package channels

import (
    "fmt"
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/PluralityNET/PluralityServer/pkg/grpc_util"
    "github.com/PluralityNET/PluralityServer/pkg/logger"
    "github.com/PluralityNET/PluralityServer/mtproto"
)

// channels.leaveChannel#f836aa95 channel:InputChannel = Updates;
func (s *ChannelsServiceImpl) ChannelsLeaveChannel(ctx context.Context, request *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("channels.leaveChannel - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // Sorry: not impl ChannelsLeaveChannel logic
    glog.Warning("channels.leaveChannel blocked, License key from https://nebula.chat required to unlock enterprise features.")

    return nil, fmt.Errorf("not imp ChannelsLeaveChannel")
}
