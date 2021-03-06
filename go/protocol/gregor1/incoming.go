// Auto-generated by avdl-compiler v1.3.19 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/gregor1/incoming.avdl

package gregor1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type SyncResult struct {
	Msgs []InBandMessage `codec:"msgs" json:"msgs"`
	Hash []byte          `codec:"hash" json:"hash"`
}

func (o SyncResult) DeepCopy() SyncResult {
	return SyncResult{
		Msgs: (func(x []InBandMessage) []InBandMessage {
			var ret []InBandMessage
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Msgs),
		Hash: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte(nil), x...)
		})(o.Hash),
	}
}

// DescribeConnectedUsers will take a list of users, and return the list of users
// which are connected to any Gregor in the cluster, and what devices (and device type)
// those users are connected with.
type ConnectedDevice struct {
	DeviceID       DeviceID `codec:"deviceID" json:"deviceID"`
	DeviceType     string   `codec:"deviceType" json:"deviceType"`
	DevicePlatform string   `codec:"devicePlatform" json:"devicePlatform"`
	UserAgent      string   `codec:"userAgent" json:"userAgent"`
}

func (o ConnectedDevice) DeepCopy() ConnectedDevice {
	return ConnectedDevice{
		DeviceID:       o.DeviceID.DeepCopy(),
		DeviceType:     o.DeviceType,
		DevicePlatform: o.DevicePlatform,
		UserAgent:      o.UserAgent,
	}
}

type ConnectedUser struct {
	Uid     UID               `codec:"uid" json:"uid"`
	Devices []ConnectedDevice `codec:"devices" json:"devices"`
}

func (o ConnectedUser) DeepCopy() ConnectedUser {
	return ConnectedUser{
		Uid: o.Uid.DeepCopy(),
		Devices: (func(x []ConnectedDevice) []ConnectedDevice {
			var ret []ConnectedDevice
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Devices),
	}
}

type SyncArg struct {
	Uid      UID      `codec:"uid" json:"uid"`
	Deviceid DeviceID `codec:"deviceid" json:"deviceid"`
	Ctime    Time     `codec:"ctime" json:"ctime"`
}

func (o SyncArg) DeepCopy() SyncArg {
	return SyncArg{
		Uid:      o.Uid.DeepCopy(),
		Deviceid: o.Deviceid.DeepCopy(),
		Ctime:    o.Ctime.DeepCopy(),
	}
}

type ConsumeMessageArg struct {
	M Message `codec:"m" json:"m"`
}

func (o ConsumeMessageArg) DeepCopy() ConsumeMessageArg {
	return ConsumeMessageArg{
		M: o.M.DeepCopy(),
	}
}

type ConsumePublishMessageArg struct {
	M Message `codec:"m" json:"m"`
}

func (o ConsumePublishMessageArg) DeepCopy() ConsumePublishMessageArg {
	return ConsumePublishMessageArg{
		M: o.M.DeepCopy(),
	}
}

type PingArg struct {
}

func (o PingArg) DeepCopy() PingArg {
	return PingArg{}
}

type VersionArg struct {
	Uid UID `codec:"uid" json:"uid"`
}

func (o VersionArg) DeepCopy() VersionArg {
	return VersionArg{
		Uid: o.Uid.DeepCopy(),
	}
}

type StateArg struct {
	Uid          UID          `codec:"uid" json:"uid"`
	Deviceid     DeviceID     `codec:"deviceid" json:"deviceid"`
	TimeOrOffset TimeOrOffset `codec:"timeOrOffset" json:"timeOrOffset"`
}

func (o StateArg) DeepCopy() StateArg {
	return StateArg{
		Uid:          o.Uid.DeepCopy(),
		Deviceid:     o.Deviceid.DeepCopy(),
		TimeOrOffset: o.TimeOrOffset.DeepCopy(),
	}
}

type StateByCategoryPrefixArg struct {
	Uid            UID          `codec:"uid" json:"uid"`
	Deviceid       DeviceID     `codec:"deviceid" json:"deviceid"`
	TimeOrOffset   TimeOrOffset `codec:"timeOrOffset" json:"timeOrOffset"`
	CategoryPrefix Category     `codec:"categoryPrefix" json:"categoryPrefix"`
}

func (o StateByCategoryPrefixArg) DeepCopy() StateByCategoryPrefixArg {
	return StateByCategoryPrefixArg{
		Uid:            o.Uid.DeepCopy(),
		Deviceid:       o.Deviceid.DeepCopy(),
		TimeOrOffset:   o.TimeOrOffset.DeepCopy(),
		CategoryPrefix: o.CategoryPrefix.DeepCopy(),
	}
}

type DescribeConnectedUsersArg struct {
	Uids []UID `codec:"uids" json:"uids"`
}

func (o DescribeConnectedUsersArg) DeepCopy() DescribeConnectedUsersArg {
	return DescribeConnectedUsersArg{
		Uids: (func(x []UID) []UID {
			var ret []UID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Uids),
	}
}

type DescribeConnectedUsersInternalArg struct {
	Uids []UID `codec:"uids" json:"uids"`
}

func (o DescribeConnectedUsersInternalArg) DeepCopy() DescribeConnectedUsersInternalArg {
	return DescribeConnectedUsersInternalArg{
		Uids: (func(x []UID) []UID {
			var ret []UID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Uids),
	}
}

type IncomingInterface interface {
	Sync(context.Context, SyncArg) (SyncResult, error)
	ConsumeMessage(context.Context, Message) error
	ConsumePublishMessage(context.Context, Message) error
	Ping(context.Context) (string, error)
	Version(context.Context, UID) (string, error)
	State(context.Context, StateArg) (State, error)
	// StateByCategoryPrefix loads the messages of the user's state whose
	// categories are prefixed by the given prefix
	StateByCategoryPrefix(context.Context, StateByCategoryPrefixArg) (State, error)
	DescribeConnectedUsers(context.Context, []UID) ([]ConnectedUser, error)
	DescribeConnectedUsersInternal(context.Context, []UID) ([]ConnectedUser, error)
}

func IncomingProtocol(i IncomingInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "gregor.1.incoming",
		Methods: map[string]rpc.ServeHandlerDescription{
			"sync": {
				MakeArg: func() interface{} {
					ret := make([]SyncArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SyncArg)
					if !ok {
						err = rpc.NewTypeError((*[]SyncArg)(nil), args)
						return
					}
					ret, err = i.Sync(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"consumeMessage": {
				MakeArg: func() interface{} {
					ret := make([]ConsumeMessageArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ConsumeMessageArg)
					if !ok {
						err = rpc.NewTypeError((*[]ConsumeMessageArg)(nil), args)
						return
					}
					err = i.ConsumeMessage(ctx, (*typedArgs)[0].M)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"consumePublishMessage": {
				MakeArg: func() interface{} {
					ret := make([]ConsumePublishMessageArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ConsumePublishMessageArg)
					if !ok {
						err = rpc.NewTypeError((*[]ConsumePublishMessageArg)(nil), args)
						return
					}
					err = i.ConsumePublishMessage(ctx, (*typedArgs)[0].M)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"ping": {
				MakeArg: func() interface{} {
					ret := make([]PingArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.Ping(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"version": {
				MakeArg: func() interface{} {
					ret := make([]VersionArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]VersionArg)
					if !ok {
						err = rpc.NewTypeError((*[]VersionArg)(nil), args)
						return
					}
					ret, err = i.Version(ctx, (*typedArgs)[0].Uid)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"state": {
				MakeArg: func() interface{} {
					ret := make([]StateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]StateArg)
					if !ok {
						err = rpc.NewTypeError((*[]StateArg)(nil), args)
						return
					}
					ret, err = i.State(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"stateByCategoryPrefix": {
				MakeArg: func() interface{} {
					ret := make([]StateByCategoryPrefixArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]StateByCategoryPrefixArg)
					if !ok {
						err = rpc.NewTypeError((*[]StateByCategoryPrefixArg)(nil), args)
						return
					}
					ret, err = i.StateByCategoryPrefix(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"describeConnectedUsers": {
				MakeArg: func() interface{} {
					ret := make([]DescribeConnectedUsersArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DescribeConnectedUsersArg)
					if !ok {
						err = rpc.NewTypeError((*[]DescribeConnectedUsersArg)(nil), args)
						return
					}
					ret, err = i.DescribeConnectedUsers(ctx, (*typedArgs)[0].Uids)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"describeConnectedUsersInternal": {
				MakeArg: func() interface{} {
					ret := make([]DescribeConnectedUsersInternalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DescribeConnectedUsersInternalArg)
					if !ok {
						err = rpc.NewTypeError((*[]DescribeConnectedUsersInternalArg)(nil), args)
						return
					}
					ret, err = i.DescribeConnectedUsersInternal(ctx, (*typedArgs)[0].Uids)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type IncomingClient struct {
	Cli rpc.GenericClient
}

func (c IncomingClient) Sync(ctx context.Context, __arg SyncArg) (res SyncResult, err error) {
	err = c.Cli.Call(ctx, "gregor.1.incoming.sync", []interface{}{__arg}, &res)
	return
}

func (c IncomingClient) ConsumeMessage(ctx context.Context, m Message) (err error) {
	__arg := ConsumeMessageArg{M: m}
	err = c.Cli.Call(ctx, "gregor.1.incoming.consumeMessage", []interface{}{__arg}, nil)
	return
}

func (c IncomingClient) ConsumePublishMessage(ctx context.Context, m Message) (err error) {
	__arg := ConsumePublishMessageArg{M: m}
	err = c.Cli.Call(ctx, "gregor.1.incoming.consumePublishMessage", []interface{}{__arg}, nil)
	return
}

func (c IncomingClient) Ping(ctx context.Context) (res string, err error) {
	err = c.Cli.Call(ctx, "gregor.1.incoming.ping", []interface{}{PingArg{}}, &res)
	return
}

func (c IncomingClient) Version(ctx context.Context, uid UID) (res string, err error) {
	__arg := VersionArg{Uid: uid}
	err = c.Cli.Call(ctx, "gregor.1.incoming.version", []interface{}{__arg}, &res)
	return
}

func (c IncomingClient) State(ctx context.Context, __arg StateArg) (res State, err error) {
	err = c.Cli.Call(ctx, "gregor.1.incoming.state", []interface{}{__arg}, &res)
	return
}

// StateByCategoryPrefix loads the messages of the user's state whose
// categories are prefixed by the given prefix
func (c IncomingClient) StateByCategoryPrefix(ctx context.Context, __arg StateByCategoryPrefixArg) (res State, err error) {
	err = c.Cli.Call(ctx, "gregor.1.incoming.stateByCategoryPrefix", []interface{}{__arg}, &res)
	return
}

func (c IncomingClient) DescribeConnectedUsers(ctx context.Context, uids []UID) (res []ConnectedUser, err error) {
	__arg := DescribeConnectedUsersArg{Uids: uids}
	err = c.Cli.Call(ctx, "gregor.1.incoming.describeConnectedUsers", []interface{}{__arg}, &res)
	return
}

func (c IncomingClient) DescribeConnectedUsersInternal(ctx context.Context, uids []UID) (res []ConnectedUser, err error) {
	__arg := DescribeConnectedUsersInternalArg{Uids: uids}
	err = c.Cli.Call(ctx, "gregor.1.incoming.describeConnectedUsersInternal", []interface{}{__arg}, &res)
	return
}
