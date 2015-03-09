package keybase_1

import (
	"github.com/maxtaco/go-framed-msgpack-rpc/rpc2"
)

type GenericClient interface {
	Call(s string, args interface{}, res interface{}) error
}

type StringKVPair struct {
	Key   string `codec:"key" json:"key"`
	Value string `codec:"value" json:"value"`
}

type Status struct {
	Code   int            `codec:"code" json:"code"`
	Name   string         `codec:"name" json:"name"`
	Desc   string         `codec:"desc" json:"desc"`
	Fields []StringKVPair `codec:"fields" json:"fields"`
}

type UID [16]byte
type LoadUserArg struct {
	Uid      *UID    `codec:"uid,omitempty" json:"uid"`
	Username *string `codec:"username,omitempty" json:"username"`
	Self     bool    `codec:"self" json:"self"`
}

type FOKID struct {
	PgpFingerprint *[]byte `codec:"pgpFingerprint,omitempty" json:"pgpFingerprint"`
	Kid            *[]byte `codec:"kid,omitempty" json:"kid"`
}

type Text struct {
	Data   string `codec:"data" json:"data"`
	Markup bool   `codec:"markup" json:"markup"`
}

type PgpIdentity struct {
	Username string `codec:"username" json:"username"`
	Comment  string `codec:"comment" json:"comment"`
	Email    string `codec:"email" json:"email"`
}

type Image struct {
	Url    string `codec:"url" json:"url"`
	Width  int    `codec:"width" json:"width"`
	Height int    `codec:"height" json:"height"`
}

type User struct {
	Uid      UID    `codec:"uid" json:"uid"`
	Username string `codec:"username" json:"username"`
	Image    *Image `codec:"image,omitempty" json:"image"`
}

type Device struct {
	Type     string `codec:"type" json:"type"`
	Name     string `codec:"name" json:"name"`
	DeviceID string `codec:"deviceID" json:"deviceID"`
}

type SIGID [32]byte
type AnnounceSessionArg struct {
	Sid string `codec:"sid" json:"sid"`
}

type GetArg struct {
	Blockid []byte `codec:"blockid" json:"blockid"`
	Uid     UID    `codec:"uid" json:"uid"`
}

type DeleteArg struct {
	Blockid []byte `codec:"blockid" json:"blockid"`
	Uid     UID    `codec:"uid" json:"uid"`
}

type PutArg struct {
	Blockid []byte `codec:"blockid" json:"blockid"`
	Uid     UID    `codec:"uid" json:"uid"`
	Buf     []byte `codec:"buf" json:"buf"`
}

type BlockInterface interface {
	AnnounceSession(string) error
	Get(GetArg) ([]byte, error)
	Delete(DeleteArg) error
	Put(PutArg) error
}

func BlockProtocol(i BlockInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.block",
		Methods: map[string]rpc2.ServeHook{
			"announceSession": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]AnnounceSessionArg, 1)
				if err = nxt(&args); err == nil {
					err = i.AnnounceSession(args[0].Sid)
				}
				return
			},
			"get": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]GetArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.Get(args[0])
				}
				return
			},
			"delete": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]DeleteArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Delete(args[0])
				}
				return
			},
			"put": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PutArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Put(args[0])
				}
				return
			},
		},
	}

}

type BlockClient struct {
	Cli GenericClient
}

func (c BlockClient) AnnounceSession(sid string) (err error) {
	__arg := AnnounceSessionArg{Sid: sid}
	err = c.Cli.Call("keybase.1.block.announceSession", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) Get(__arg GetArg) (res []byte, err error) {
	err = c.Cli.Call("keybase.1.block.get", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) Delete(__arg DeleteArg) (err error) {
	err = c.Cli.Call("keybase.1.block.delete", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) Put(__arg PutArg) (err error) {
	err = c.Cli.Call("keybase.1.block.put", []interface{}{__arg}, nil)
	return
}

type AvdlFile struct {
	Id int `codec:"id" json:"id"`
}

type GetCurrentStatusRes struct {
	Configured bool  `codec:"configured" json:"configured"`
	Registered bool  `codec:"registered" json:"registered"`
	LoggedIn   bool  `codec:"loggedIn" json:"loggedIn"`
	User       *User `codec:"user,omitempty" json:"user"`
}

type Config struct {
	ServerURI  string `codec:"serverURI" json:"serverURI"`
	SocketFile string `codec:"socketFile" json:"socketFile"`
	GpgExists  bool   `codec:"gpgExists" json:"gpgExists"`
	GpgPath    string `codec:"gpgPath" json:"gpgPath"`
}

type GetCurrentStatusArg struct {
}

type GetConfigArg struct {
}

type ConfigInterface interface {
	GetCurrentStatus() (GetCurrentStatusRes, error)
	GetConfig() (Config, error)
}

func ConfigProtocol(i ConfigInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.config",
		Methods: map[string]rpc2.ServeHook{
			"getCurrentStatus": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]GetCurrentStatusArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.GetCurrentStatus()
				}
				return
			},
			"getConfig": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]GetConfigArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.GetConfig()
				}
				return
			},
		},
	}

}

type ConfigClient struct {
	Cli GenericClient
}

func (c ConfigClient) GetCurrentStatus() (res GetCurrentStatusRes, err error) {
	err = c.Cli.Call("keybase.1.config.getCurrentStatus", []interface{}{GetCurrentStatusArg{}}, &res)
	return
}

func (c ConfigClient) GetConfig() (res Config, err error) {
	err = c.Cli.Call("keybase.1.config.getConfig", []interface{}{GetConfigArg{}}, &res)
	return
}

type DeviceListArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type DeviceInterface interface {
	DeviceList(int) ([]Device, error)
}

func DeviceProtocol(i DeviceInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.device",
		Methods: map[string]rpc2.ServeHook{
			"deviceList": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]DeviceListArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.DeviceList(args[0].SessionID)
				}
				return
			},
		},
	}

}

type DeviceClient struct {
	Cli GenericClient
}

func (c DeviceClient) DeviceList(sessionID int) (res []Device, err error) {
	__arg := DeviceListArg{SessionID: sessionID}
	err = c.Cli.Call("keybase.1.device.deviceList", []interface{}{__arg}, &res)
	return
}

type DeviceSignerKind int

const (
	DeviceSignerKind_DEVICE = 0
	DeviceSignerKind_PGP    = 1
)

type SelectSignerAction int

const (
	SelectSignerAction_SIGN          = 0
	SelectSignerAction_CANCEL        = 1
	SelectSignerAction_RESET_ACCOUNT = 2
)

type DeviceSigner struct {
	Kind     DeviceSignerKind `codec:"kind" json:"kind"`
	DeviceID *string          `codec:"deviceID,omitempty" json:"deviceID"`
}

type SelectSignerRes struct {
	Action SelectSignerAction `codec:"action" json:"action"`
	Signer *DeviceSigner      `codec:"signer,omitempty" json:"signer"`
}

type PromptDeviceNameArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type SelectSignerArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Devices   []Device `codec:"devices" json:"devices"`
	HasPGP    bool     `codec:"hasPGP" json:"hasPGP"`
}

type DisplaySecretWordsArg struct {
	SessionID       int    `codec:"sessionID" json:"sessionID"`
	Secret          string `codec:"secret" json:"secret"`
	XDevDescription string `codec:"xDevDescription" json:"xDevDescription"`
}

type DoctorUiInterface interface {
	PromptDeviceName(int) (string, error)
	SelectSigner(SelectSignerArg) (SelectSignerRes, error)
	DisplaySecretWords(DisplaySecretWordsArg) error
}

func DoctorUiProtocol(i DoctorUiInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.doctorUi",
		Methods: map[string]rpc2.ServeHook{
			"promptDeviceName": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PromptDeviceNameArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.PromptDeviceName(args[0].SessionID)
				}
				return
			},
			"selectSigner": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SelectSignerArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.SelectSigner(args[0])
				}
				return
			},
			"displaySecretWords": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]DisplaySecretWordsArg, 1)
				if err = nxt(&args); err == nil {
					err = i.DisplaySecretWords(args[0])
				}
				return
			},
		},
	}

}

type DoctorUiClient struct {
	Cli GenericClient
}

func (c DoctorUiClient) PromptDeviceName(sessionID int) (res string, err error) {
	__arg := PromptDeviceNameArg{SessionID: sessionID}
	err = c.Cli.Call("keybase.1.doctorUi.promptDeviceName", []interface{}{__arg}, &res)
	return
}

func (c DoctorUiClient) SelectSigner(__arg SelectSignerArg) (res SelectSignerRes, err error) {
	err = c.Cli.Call("keybase.1.doctorUi.selectSigner", []interface{}{__arg}, &res)
	return
}

func (c DoctorUiClient) DisplaySecretWords(__arg DisplaySecretWordsArg) (err error) {
	err = c.Cli.Call("keybase.1.doctorUi.displaySecretWords", []interface{}{__arg}, nil)
	return
}

type AddGpgKeyArg struct {
}

type GpgInterface interface {
	AddGpgKey() error
}

func GpgProtocol(i GpgInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.gpg",
		Methods: map[string]rpc2.ServeHook{
			"addGpgKey": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]AddGpgKeyArg, 1)
				if err = nxt(&args); err == nil {
					err = i.AddGpgKey()
				}
				return
			},
		},
	}

}

type GpgClient struct {
	Cli GenericClient
}

func (c GpgClient) AddGpgKey() (err error) {
	err = c.Cli.Call("keybase.1.gpg.addGpgKey", []interface{}{AddGpgKeyArg{}}, nil)
	return
}

type GPGKey struct {
	Algorithm  string   `codec:"algorithm" json:"algorithm"`
	KeyID      string   `codec:"keyID" json:"keyID"`
	Creation   string   `codec:"creation" json:"creation"`
	Expiration string   `codec:"expiration" json:"expiration"`
	Identities []string `codec:"identities" json:"identities"`
}

type SelectKeyRes struct {
	KeyID        string `codec:"keyID" json:"keyID"`
	DoSecretPush bool   `codec:"doSecretPush" json:"doSecretPush"`
}

type WantToAddGPGKeyArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type SelectKeyAndPushOptionArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Keys      []GPGKey `codec:"keys" json:"keys"`
}

type SelectKeyArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Keys      []GPGKey `codec:"keys" json:"keys"`
}

type GpgUiInterface interface {
	WantToAddGPGKey(int) (bool, error)
	SelectKeyAndPushOption(SelectKeyAndPushOptionArg) (SelectKeyRes, error)
	SelectKey(SelectKeyArg) (string, error)
}

func GpgUiProtocol(i GpgUiInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.gpgUi",
		Methods: map[string]rpc2.ServeHook{
			"wantToAddGPGKey": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]WantToAddGPGKeyArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.WantToAddGPGKey(args[0].SessionID)
				}
				return
			},
			"selectKeyAndPushOption": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SelectKeyAndPushOptionArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.SelectKeyAndPushOption(args[0])
				}
				return
			},
			"selectKey": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SelectKeyArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.SelectKey(args[0])
				}
				return
			},
		},
	}

}

type GpgUiClient struct {
	Cli GenericClient
}

func (c GpgUiClient) WantToAddGPGKey(sessionID int) (res bool, err error) {
	__arg := WantToAddGPGKeyArg{SessionID: sessionID}
	err = c.Cli.Call("keybase.1.gpgUi.wantToAddGPGKey", []interface{}{__arg}, &res)
	return
}

func (c GpgUiClient) SelectKeyAndPushOption(__arg SelectKeyAndPushOptionArg) (res SelectKeyRes, err error) {
	err = c.Cli.Call("keybase.1.gpgUi.selectKeyAndPushOption", []interface{}{__arg}, &res)
	return
}

func (c GpgUiClient) SelectKey(__arg SelectKeyArg) (res string, err error) {
	err = c.Cli.Call("keybase.1.gpgUi.selectKey", []interface{}{__arg}, &res)
	return
}

type TrackDiffType int

const (
	TrackDiffType_NONE           = 0
	TrackDiffType_ERROR          = 1
	TrackDiffType_CLASH          = 2
	TrackDiffType_DELETED        = 3
	TrackDiffType_UPGRADED       = 4
	TrackDiffType_NEW            = 5
	TrackDiffType_REMOTE_FAIL    = 6
	TrackDiffType_REMOTE_WORKING = 7
	TrackDiffType_REMOTE_CHANGED = 8
)

type TrackDiff struct {
	Type          TrackDiffType `codec:"type" json:"type"`
	DisplayMarkup string        `codec:"displayMarkup" json:"displayMarkup"`
}

type TrackSummary struct {
	Time     int  `codec:"time" json:"time"`
	IsRemote bool `codec:"isRemote" json:"isRemote"`
}

type IdentifyOutcome struct {
	Status            *Status       `codec:"status,omitempty" json:"status"`
	Warnings          []string      `codec:"warnings" json:"warnings"`
	TrackUsed         *TrackSummary `codec:"trackUsed,omitempty" json:"trackUsed"`
	NumTrackFailures  int           `codec:"numTrackFailures" json:"numTrackFailures"`
	NumTrackChanges   int           `codec:"numTrackChanges" json:"numTrackChanges"`
	NumProofFailures  int           `codec:"numProofFailures" json:"numProofFailures"`
	NumDeleted        int           `codec:"numDeleted" json:"numDeleted"`
	NumProofSuccesses int           `codec:"numProofSuccesses" json:"numProofSuccesses"`
	Deleted           []TrackDiff   `codec:"deleted" json:"deleted"`
}

type IdentifyRes struct {
	User    *User           `codec:"user,omitempty" json:"user"`
	Outcome IdentifyOutcome `codec:"outcome" json:"outcome"`
}

type IdentifyArg struct {
	SessionID      int    `codec:"sessionID" json:"sessionID"`
	Uid            UID    `codec:"uid" json:"uid"`
	Username       string `codec:"username" json:"username"`
	TrackStatement bool   `codec:"trackStatement" json:"trackStatement"`
	Luba           bool   `codec:"luba" json:"luba"`
	LoadSelf       bool   `codec:"loadSelf" json:"loadSelf"`
}

type IdentifyDefaultArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Username  string `codec:"username" json:"username"`
}

type IdentifyInterface interface {
	Identify(IdentifyArg) (IdentifyRes, error)
	IdentifyDefault(IdentifyDefaultArg) (IdentifyRes, error)
}

func IdentifyProtocol(i IdentifyInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.identify",
		Methods: map[string]rpc2.ServeHook{
			"identify": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]IdentifyArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.Identify(args[0])
				}
				return
			},
			"identifyDefault": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]IdentifyDefaultArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.IdentifyDefault(args[0])
				}
				return
			},
		},
	}

}

type IdentifyClient struct {
	Cli GenericClient
}

func (c IdentifyClient) Identify(__arg IdentifyArg) (res IdentifyRes, err error) {
	err = c.Cli.Call("keybase.1.identify.identify", []interface{}{__arg}, &res)
	return
}

func (c IdentifyClient) IdentifyDefault(__arg IdentifyDefaultArg) (res IdentifyRes, err error) {
	err = c.Cli.Call("keybase.1.identify.identifyDefault", []interface{}{__arg}, &res)
	return
}

type ProofStatus struct {
	State  int    `codec:"state" json:"state"`
	Status int    `codec:"status" json:"status"`
	Desc   string `codec:"desc" json:"desc"`
}

type RemoteProof struct {
	ProofType     int    `codec:"proofType" json:"proofType"`
	Key           string `codec:"key" json:"key"`
	Value         string `codec:"value" json:"value"`
	DisplayMarkup string `codec:"displayMarkup" json:"displayMarkup"`
	SigId         SIGID  `codec:"sigId" json:"sigId"`
	Mtime         int    `codec:"mtime" json:"mtime"`
}

type IdentifyRow struct {
	RowId     int         `codec:"rowId" json:"rowId"`
	Proof     RemoteProof `codec:"proof" json:"proof"`
	TrackDiff *TrackDiff  `codec:"trackDiff,omitempty" json:"trackDiff"`
}

type IdentifyKey struct {
	PgpFingerprint []byte     `codec:"pgpFingerprint" json:"pgpFingerprint"`
	KID            []byte     `codec:"KID" json:"KID"`
	TrackDiff      *TrackDiff `codec:"trackDiff,omitempty" json:"trackDiff"`
}

type Cryptocurrency struct {
	RowId   int    `codec:"rowId" json:"rowId"`
	Pkhash  []byte `codec:"pkhash" json:"pkhash"`
	Address string `codec:"address" json:"address"`
}

type Identity struct {
	Status          *Status          `codec:"status,omitempty" json:"status"`
	WhenLastTracked int              `codec:"whenLastTracked" json:"whenLastTracked"`
	Key             IdentifyKey      `codec:"key" json:"key"`
	Proofs          []IdentifyRow    `codec:"proofs" json:"proofs"`
	Cryptocurrency  []Cryptocurrency `codec:"cryptocurrency" json:"cryptocurrency"`
	Deleted         []TrackDiff      `codec:"deleted" json:"deleted"`
}

type SigHint struct {
	RemoteId  string `codec:"remoteId" json:"remoteId"`
	HumanUrl  string `codec:"humanUrl" json:"humanUrl"`
	ApiUrl    string `codec:"apiUrl" json:"apiUrl"`
	CheckText string `codec:"checkText" json:"checkText"`
}

type CheckResult struct {
	ProofStatus   ProofStatus `codec:"proofStatus" json:"proofStatus"`
	Timestamp     int         `codec:"timestamp" json:"timestamp"`
	DisplayMarkup string      `codec:"displayMarkup" json:"displayMarkup"`
}

type LinkCheckResult struct {
	ProofId     int          `codec:"proofId" json:"proofId"`
	ProofStatus ProofStatus  `codec:"proofStatus" json:"proofStatus"`
	Cached      *CheckResult `codec:"cached,omitempty" json:"cached"`
	Diff        *TrackDiff   `codec:"diff,omitempty" json:"diff"`
	RemoteDiff  *TrackDiff   `codec:"remoteDiff,omitempty" json:"remoteDiff"`
	Hint        *SigHint     `codec:"hint,omitempty" json:"hint"`
}

type FinishAndPromptRes struct {
	TrackLocal  bool `codec:"trackLocal" json:"trackLocal"`
	TrackRemote bool `codec:"trackRemote" json:"trackRemote"`
}

type FinishAndPromptArg struct {
	SessionID int             `codec:"sessionID" json:"sessionID"`
	Outcome   IdentifyOutcome `codec:"outcome" json:"outcome"`
}

type FinishWebProofCheckArg struct {
	SessionID int             `codec:"sessionID" json:"sessionID"`
	Rp        RemoteProof     `codec:"rp" json:"rp"`
	Lcr       LinkCheckResult `codec:"lcr" json:"lcr"`
}

type FinishSocialProofCheckArg struct {
	SessionID int             `codec:"sessionID" json:"sessionID"`
	Rp        RemoteProof     `codec:"rp" json:"rp"`
	Lcr       LinkCheckResult `codec:"lcr" json:"lcr"`
}

type DisplayCryptocurrencyArg struct {
	SessionID int            `codec:"sessionID" json:"sessionID"`
	C         Cryptocurrency `codec:"c" json:"c"`
}

type DisplayKeyArg struct {
	SessionID int        `codec:"sessionID" json:"sessionID"`
	Fokid     FOKID      `codec:"fokid" json:"fokid"`
	Diff      *TrackDiff `codec:"diff,omitempty" json:"diff"`
}

type ReportLastTrackArg struct {
	SessionID int           `codec:"sessionID" json:"sessionID"`
	Track     *TrackSummary `codec:"track,omitempty" json:"track"`
}

type LaunchNetworkChecksArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Id        Identity `codec:"id" json:"id"`
	User      User     `codec:"user" json:"user"`
}

type DisplayTrackStatementArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Stmt      string `codec:"stmt" json:"stmt"`
}

type IdentifyUiInterface interface {
	FinishAndPrompt(FinishAndPromptArg) (FinishAndPromptRes, error)
	FinishWebProofCheck(FinishWebProofCheckArg) error
	FinishSocialProofCheck(FinishSocialProofCheckArg) error
	DisplayCryptocurrency(DisplayCryptocurrencyArg) error
	DisplayKey(DisplayKeyArg) error
	ReportLastTrack(ReportLastTrackArg) error
	LaunchNetworkChecks(LaunchNetworkChecksArg) error
	DisplayTrackStatement(DisplayTrackStatementArg) error
}

func IdentifyUiProtocol(i IdentifyUiInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.identifyUi",
		Methods: map[string]rpc2.ServeHook{
			"finishAndPrompt": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]FinishAndPromptArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.FinishAndPrompt(args[0])
				}
				return
			},
			"finishWebProofCheck": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]FinishWebProofCheckArg, 1)
				if err = nxt(&args); err == nil {
					err = i.FinishWebProofCheck(args[0])
				}
				return
			},
			"finishSocialProofCheck": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]FinishSocialProofCheckArg, 1)
				if err = nxt(&args); err == nil {
					err = i.FinishSocialProofCheck(args[0])
				}
				return
			},
			"displayCryptocurrency": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]DisplayCryptocurrencyArg, 1)
				if err = nxt(&args); err == nil {
					err = i.DisplayCryptocurrency(args[0])
				}
				return
			},
			"displayKey": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]DisplayKeyArg, 1)
				if err = nxt(&args); err == nil {
					err = i.DisplayKey(args[0])
				}
				return
			},
			"reportLastTrack": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]ReportLastTrackArg, 1)
				if err = nxt(&args); err == nil {
					err = i.ReportLastTrack(args[0])
				}
				return
			},
			"launchNetworkChecks": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]LaunchNetworkChecksArg, 1)
				if err = nxt(&args); err == nil {
					err = i.LaunchNetworkChecks(args[0])
				}
				return
			},
			"displayTrackStatement": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]DisplayTrackStatementArg, 1)
				if err = nxt(&args); err == nil {
					err = i.DisplayTrackStatement(args[0])
				}
				return
			},
		},
	}

}

type IdentifyUiClient struct {
	Cli GenericClient
}

func (c IdentifyUiClient) FinishAndPrompt(__arg FinishAndPromptArg) (res FinishAndPromptRes, err error) {
	err = c.Cli.Call("keybase.1.identifyUi.finishAndPrompt", []interface{}{__arg}, &res)
	return
}

func (c IdentifyUiClient) FinishWebProofCheck(__arg FinishWebProofCheckArg) (err error) {
	err = c.Cli.Call("keybase.1.identifyUi.finishWebProofCheck", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) FinishSocialProofCheck(__arg FinishSocialProofCheckArg) (err error) {
	err = c.Cli.Call("keybase.1.identifyUi.finishSocialProofCheck", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) DisplayCryptocurrency(__arg DisplayCryptocurrencyArg) (err error) {
	err = c.Cli.Call("keybase.1.identifyUi.displayCryptocurrency", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) DisplayKey(__arg DisplayKeyArg) (err error) {
	err = c.Cli.Call("keybase.1.identifyUi.displayKey", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) ReportLastTrack(__arg ReportLastTrackArg) (err error) {
	err = c.Cli.Call("keybase.1.identifyUi.reportLastTrack", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) LaunchNetworkChecks(__arg LaunchNetworkChecksArg) (err error) {
	err = c.Cli.Call("keybase.1.identifyUi.launchNetworkChecks", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) DisplayTrackStatement(__arg DisplayTrackStatementArg) (err error) {
	err = c.Cli.Call("keybase.1.identifyUi.displayTrackStatement", []interface{}{__arg}, nil)
	return
}

type LogLevel int

const (
	LogLevel_NONE     = 0
	LogLevel_DEBUG    = 1
	LogLevel_INFO     = 2
	LogLevel_NOTICE   = 3
	LogLevel_WARN     = 4
	LogLevel_ERROR    = 5
	LogLevel_CRITICAL = 6
)

type LogArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Level     LogLevel `codec:"level" json:"level"`
	Text      Text     `codec:"text" json:"text"`
}

type LogUiInterface interface {
	Log(LogArg) error
}

func LogUiProtocol(i LogUiInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.logUi",
		Methods: map[string]rpc2.ServeHook{
			"log": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]LogArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Log(args[0])
				}
				return
			},
		},
	}

}

type LogUiClient struct {
	Cli GenericClient
}

func (c LogUiClient) Log(__arg LogArg) (err error) {
	err = c.Cli.Call("keybase.1.logUi.log", []interface{}{__arg}, nil)
	return
}

type PassphraseLoginArg struct {
	Identify   bool   `codec:"identify" json:"identify"`
	Username   string `codec:"username" json:"username"`
	Passphrase string `codec:"passphrase" json:"passphrase"`
}

type PubkeyLoginArg struct {
}

type LogoutArg struct {
}

type SwitchUserArg struct {
	Username string `codec:"username" json:"username"`
}

type LoginInterface interface {
	PassphraseLogin(PassphraseLoginArg) error
	PubkeyLogin() error
	Logout() error
	SwitchUser(string) error
}

func LoginProtocol(i LoginInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.login",
		Methods: map[string]rpc2.ServeHook{
			"passphraseLogin": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PassphraseLoginArg, 1)
				if err = nxt(&args); err == nil {
					err = i.PassphraseLogin(args[0])
				}
				return
			},
			"pubkeyLogin": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PubkeyLoginArg, 1)
				if err = nxt(&args); err == nil {
					err = i.PubkeyLogin()
				}
				return
			},
			"logout": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]LogoutArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Logout()
				}
				return
			},
			"switchUser": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SwitchUserArg, 1)
				if err = nxt(&args); err == nil {
					err = i.SwitchUser(args[0].Username)
				}
				return
			},
		},
	}

}

type LoginClient struct {
	Cli GenericClient
}

func (c LoginClient) PassphraseLogin(__arg PassphraseLoginArg) (err error) {
	err = c.Cli.Call("keybase.1.login.passphraseLogin", []interface{}{__arg}, nil)
	return
}

func (c LoginClient) PubkeyLogin() (err error) {
	err = c.Cli.Call("keybase.1.login.pubkeyLogin", []interface{}{PubkeyLoginArg{}}, nil)
	return
}

func (c LoginClient) Logout() (err error) {
	err = c.Cli.Call("keybase.1.login.logout", []interface{}{LogoutArg{}}, nil)
	return
}

func (c LoginClient) SwitchUser(username string) (err error) {
	__arg := SwitchUserArg{Username: username}
	err = c.Cli.Call("keybase.1.login.switchUser", []interface{}{__arg}, nil)
	return
}

type GetEmailOrUsernameArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type LoginUiInterface interface {
	GetEmailOrUsername(int) (string, error)
}

func LoginUiProtocol(i LoginUiInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.loginUi",
		Methods: map[string]rpc2.ServeHook{
			"getEmailOrUsername": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]GetEmailOrUsernameArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.GetEmailOrUsername(args[0].SessionID)
				}
				return
			},
		},
	}

}

type LoginUiClient struct {
	Cli GenericClient
}

func (c LoginUiClient) GetEmailOrUsername(sessionID int) (res string, err error) {
	__arg := GetEmailOrUsernameArg{SessionID: sessionID}
	err = c.Cli.Call("keybase.1.loginUi.getEmailOrUsername", []interface{}{__arg}, &res)
	return
}

type Stream struct {
	Fd int `codec:"fd" json:"fd"`
}

type PgpCreateUids struct {
	UseDefault bool          `codec:"useDefault" json:"useDefault"`
	Ids        []PgpIdentity `codec:"ids" json:"ids"`
}

type PgpKeyGenArg struct {
	PrimaryBits int           `codec:"primaryBits" json:"primaryBits"`
	SubkeyBits  int           `codec:"subkeyBits" json:"subkeyBits"`
	CreateUids  PgpCreateUids `codec:"createUids" json:"createUids"`
	AllowMulti  bool          `codec:"allowMulti" json:"allowMulti"`
	DoExport    bool          `codec:"doExport" json:"doExport"`
}

type PgpKeyGenDefaultArg struct {
	CreateUids PgpCreateUids `codec:"createUids" json:"createUids"`
}

type DeletePrimaryArg struct {
}

type ShowArg struct {
}

type SelectArg struct {
	Query      string `codec:"query" json:"query"`
	AllowMulti bool   `codec:"allowMulti" json:"allowMulti"`
	SkipImport bool   `codec:"skipImport" json:"skipImport"`
}

type SaveArmoredPGPKeyArg struct {
	Key         string `codec:"key" json:"key"`
	PushPublic  bool   `codec:"pushPublic" json:"pushPublic"`
	PushPrivate bool   `codec:"pushPrivate" json:"pushPrivate"`
}

type SavePGPKeyArg struct {
	Key         []byte `codec:"key" json:"key"`
	PushPublic  bool   `codec:"pushPublic" json:"pushPublic"`
	PushPrivate bool   `codec:"pushPrivate" json:"pushPrivate"`
}

type MykeyInterface interface {
	PgpKeyGen(PgpKeyGenArg) error
	PgpKeyGenDefault(PgpCreateUids) error
	DeletePrimary() error
	Show() error
	Select(SelectArg) error
	SaveArmoredPGPKey(SaveArmoredPGPKeyArg) error
	SavePGPKey(SavePGPKeyArg) error
}

func MykeyProtocol(i MykeyInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.mykey",
		Methods: map[string]rpc2.ServeHook{
			"PgpKeyGen": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PgpKeyGenArg, 1)
				if err = nxt(&args); err == nil {
					err = i.PgpKeyGen(args[0])
				}
				return
			},
			"pgpKeyGenDefault": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PgpKeyGenDefaultArg, 1)
				if err = nxt(&args); err == nil {
					err = i.PgpKeyGenDefault(args[0].CreateUids)
				}
				return
			},
			"deletePrimary": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]DeletePrimaryArg, 1)
				if err = nxt(&args); err == nil {
					err = i.DeletePrimary()
				}
				return
			},
			"show": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]ShowArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Show()
				}
				return
			},
			"select": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SelectArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Select(args[0])
				}
				return
			},
			"saveArmoredPGPKey": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SaveArmoredPGPKeyArg, 1)
				if err = nxt(&args); err == nil {
					err = i.SaveArmoredPGPKey(args[0])
				}
				return
			},
			"savePGPKey": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SavePGPKeyArg, 1)
				if err = nxt(&args); err == nil {
					err = i.SavePGPKey(args[0])
				}
				return
			},
		},
	}

}

type MykeyClient struct {
	Cli GenericClient
}

func (c MykeyClient) PgpKeyGen(__arg PgpKeyGenArg) (err error) {
	err = c.Cli.Call("keybase.1.mykey.PgpKeyGen", []interface{}{__arg}, nil)
	return
}

func (c MykeyClient) PgpKeyGenDefault(createUids PgpCreateUids) (err error) {
	__arg := PgpKeyGenDefaultArg{CreateUids: createUids}
	err = c.Cli.Call("keybase.1.mykey.pgpKeyGenDefault", []interface{}{__arg}, nil)
	return
}

func (c MykeyClient) DeletePrimary() (err error) {
	err = c.Cli.Call("keybase.1.mykey.deletePrimary", []interface{}{DeletePrimaryArg{}}, nil)
	return
}

func (c MykeyClient) Show() (err error) {
	err = c.Cli.Call("keybase.1.mykey.show", []interface{}{ShowArg{}}, nil)
	return
}

func (c MykeyClient) Select(__arg SelectArg) (err error) {
	err = c.Cli.Call("keybase.1.mykey.select", []interface{}{__arg}, nil)
	return
}

func (c MykeyClient) SaveArmoredPGPKey(__arg SaveArmoredPGPKeyArg) (err error) {
	err = c.Cli.Call("keybase.1.mykey.saveArmoredPGPKey", []interface{}{__arg}, nil)
	return
}

func (c MykeyClient) SavePGPKey(__arg SavePGPKeyArg) (err error) {
	err = c.Cli.Call("keybase.1.mykey.savePGPKey", []interface{}{__arg}, nil)
	return
}

type PgpSignArg struct {
	Source   Stream `codec:"source" json:"source"`
	Sink     Stream `codec:"sink" json:"sink"`
	KeyQuery string `codec:"keyQuery" json:"keyQuery"`
}

type PgpcmdsInterface interface {
	PgpSign(PgpSignArg) error
}

func PgpcmdsProtocol(i PgpcmdsInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.pgpcmds",
		Methods: map[string]rpc2.ServeHook{
			"pgpSign": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PgpSignArg, 1)
				if err = nxt(&args); err == nil {
					err = i.PgpSign(args[0])
				}
				return
			},
		},
	}

}

type PgpcmdsClient struct {
	Cli GenericClient
}

func (c PgpcmdsClient) PgpSign(__arg PgpSignArg) (err error) {
	err = c.Cli.Call("keybase.1.pgpcmds.pgpSign", []interface{}{__arg}, nil)
	return
}

type ProveArg struct {
	Service  string `codec:"service" json:"service"`
	Username string `codec:"username" json:"username"`
	Force    bool   `codec:"force" json:"force"`
}

type ProveInterface interface {
	Prove(ProveArg) error
}

func ProveProtocol(i ProveInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.prove",
		Methods: map[string]rpc2.ServeHook{
			"prove": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]ProveArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Prove(args[0])
				}
				return
			},
		},
	}

}

type ProveClient struct {
	Cli GenericClient
}

func (c ProveClient) Prove(__arg ProveArg) (err error) {
	err = c.Cli.Call("keybase.1.prove.prove", []interface{}{__arg}, nil)
	return
}

type PromptOverwriteType int

const (
	PromptOverwriteType_SOCIAL = 0
	PromptOverwriteType_SITE   = 1
)

type PromptOverwriteArg struct {
	SessionID int                 `codec:"sessionID" json:"sessionID"`
	Account   string              `codec:"account" json:"account"`
	Typ       PromptOverwriteType `codec:"typ" json:"typ"`
}

type PromptUsernameArg struct {
	SessionID int     `codec:"sessionID" json:"sessionID"`
	Prompt    string  `codec:"prompt" json:"prompt"`
	PrevError *Status `codec:"prevError,omitempty" json:"prevError"`
}

type OutputPrechecksArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	Text      Text `codec:"text" json:"text"`
}

type PreProofWarningArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	Text      Text `codec:"text" json:"text"`
}

type OutputInstructionsArg struct {
	SessionID    int    `codec:"sessionID" json:"sessionID"`
	Instructions Text   `codec:"instructions" json:"instructions"`
	Proof        string `codec:"proof" json:"proof"`
}

type OkToCheckArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
	Attempt   int    `codec:"attempt" json:"attempt"`
}

type DisplayRecheckWarningArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	Text      Text `codec:"text" json:"text"`
}

type ProveUiInterface interface {
	PromptOverwrite(PromptOverwriteArg) (bool, error)
	PromptUsername(PromptUsernameArg) (string, error)
	OutputPrechecks(OutputPrechecksArg) error
	PreProofWarning(PreProofWarningArg) (bool, error)
	OutputInstructions(OutputInstructionsArg) error
	OkToCheck(OkToCheckArg) (bool, error)
	DisplayRecheckWarning(DisplayRecheckWarningArg) error
}

func ProveUiProtocol(i ProveUiInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.proveUi",
		Methods: map[string]rpc2.ServeHook{
			"promptOverwrite": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PromptOverwriteArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.PromptOverwrite(args[0])
				}
				return
			},
			"promptUsername": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PromptUsernameArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.PromptUsername(args[0])
				}
				return
			},
			"outputPrechecks": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]OutputPrechecksArg, 1)
				if err = nxt(&args); err == nil {
					err = i.OutputPrechecks(args[0])
				}
				return
			},
			"preProofWarning": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PreProofWarningArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.PreProofWarning(args[0])
				}
				return
			},
			"outputInstructions": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]OutputInstructionsArg, 1)
				if err = nxt(&args); err == nil {
					err = i.OutputInstructions(args[0])
				}
				return
			},
			"okToCheck": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]OkToCheckArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.OkToCheck(args[0])
				}
				return
			},
			"displayRecheckWarning": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]DisplayRecheckWarningArg, 1)
				if err = nxt(&args); err == nil {
					err = i.DisplayRecheckWarning(args[0])
				}
				return
			},
		},
	}

}

type ProveUiClient struct {
	Cli GenericClient
}

func (c ProveUiClient) PromptOverwrite(__arg PromptOverwriteArg) (res bool, err error) {
	err = c.Cli.Call("keybase.1.proveUi.promptOverwrite", []interface{}{__arg}, &res)
	return
}

func (c ProveUiClient) PromptUsername(__arg PromptUsernameArg) (res string, err error) {
	err = c.Cli.Call("keybase.1.proveUi.promptUsername", []interface{}{__arg}, &res)
	return
}

func (c ProveUiClient) OutputPrechecks(__arg OutputPrechecksArg) (err error) {
	err = c.Cli.Call("keybase.1.proveUi.outputPrechecks", []interface{}{__arg}, nil)
	return
}

func (c ProveUiClient) PreProofWarning(__arg PreProofWarningArg) (res bool, err error) {
	err = c.Cli.Call("keybase.1.proveUi.preProofWarning", []interface{}{__arg}, &res)
	return
}

func (c ProveUiClient) OutputInstructions(__arg OutputInstructionsArg) (err error) {
	err = c.Cli.Call("keybase.1.proveUi.outputInstructions", []interface{}{__arg}, nil)
	return
}

func (c ProveUiClient) OkToCheck(__arg OkToCheckArg) (res bool, err error) {
	err = c.Cli.Call("keybase.1.proveUi.okToCheck", []interface{}{__arg}, &res)
	return
}

func (c ProveUiClient) DisplayRecheckWarning(__arg DisplayRecheckWarningArg) (err error) {
	err = c.Cli.Call("keybase.1.proveUi.displayRecheckWarning", []interface{}{__arg}, nil)
	return
}

type SessionToken struct {
	Uid       UID    `codec:"uid" json:"uid"`
	Sid       string `codec:"sid" json:"sid"`
	Generated int    `codec:"generated" json:"generated"`
	Lifetime  int    `codec:"lifetime" json:"lifetime"`
}

type VerifySessionArg struct {
	Session string `codec:"session" json:"session"`
}

type QuotaInterface interface {
	VerifySession(string) (SessionToken, error)
}

func QuotaProtocol(i QuotaInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.quota",
		Methods: map[string]rpc2.ServeHook{
			"verifySession": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]VerifySessionArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.VerifySession(args[0].Session)
				}
				return
			},
		},
	}

}

type QuotaClient struct {
	Cli GenericClient
}

func (c QuotaClient) VerifySession(session string) (res SessionToken, err error) {
	__arg := VerifySessionArg{Session: session}
	err = c.Cli.Call("keybase.1.quota.verifySession", []interface{}{__arg}, &res)
	return
}

type SecretEntryArg struct {
	Desc   string `codec:"desc" json:"desc"`
	Prompt string `codec:"prompt" json:"prompt"`
	Err    string `codec:"err" json:"err"`
	Cancel string `codec:"cancel" json:"cancel"`
	Ok     string `codec:"ok" json:"ok"`
}

type SecretEntryRes struct {
	Text     string `codec:"text" json:"text"`
	Canceled bool   `codec:"canceled" json:"canceled"`
}

type GetSecretArg struct {
	SessionID int             `codec:"sessionID" json:"sessionID"`
	Pinentry  SecretEntryArg  `codec:"pinentry" json:"pinentry"`
	Terminal  *SecretEntryArg `codec:"terminal,omitempty" json:"terminal"`
}

type GetNewPassphraseArg struct {
	TerminalPrompt string `codec:"terminalPrompt" json:"terminalPrompt"`
	PinentryDesc   string `codec:"pinentryDesc" json:"pinentryDesc"`
	PinentryPrompt string `codec:"pinentryPrompt" json:"pinentryPrompt"`
	RetryMessage   string `codec:"retryMessage" json:"retryMessage"`
}

type GetKeybasePassphraseArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Username  string `codec:"username" json:"username"`
	Retry     string `codec:"retry" json:"retry"`
}

type SecretUiInterface interface {
	GetSecret(GetSecretArg) (SecretEntryRes, error)
	GetNewPassphrase(GetNewPassphraseArg) (string, error)
	GetKeybasePassphrase(GetKeybasePassphraseArg) (string, error)
}

func SecretUiProtocol(i SecretUiInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.secretUi",
		Methods: map[string]rpc2.ServeHook{
			"getSecret": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]GetSecretArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.GetSecret(args[0])
				}
				return
			},
			"getNewPassphrase": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]GetNewPassphraseArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.GetNewPassphrase(args[0])
				}
				return
			},
			"getKeybasePassphrase": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]GetKeybasePassphraseArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.GetKeybasePassphrase(args[0])
				}
				return
			},
		},
	}

}

type SecretUiClient struct {
	Cli GenericClient
}

func (c SecretUiClient) GetSecret(__arg GetSecretArg) (res SecretEntryRes, err error) {
	err = c.Cli.Call("keybase.1.secretUi.getSecret", []interface{}{__arg}, &res)
	return
}

func (c SecretUiClient) GetNewPassphrase(__arg GetNewPassphraseArg) (res string, err error) {
	err = c.Cli.Call("keybase.1.secretUi.getNewPassphrase", []interface{}{__arg}, &res)
	return
}

func (c SecretUiClient) GetKeybasePassphrase(__arg GetKeybasePassphraseArg) (res string, err error) {
	err = c.Cli.Call("keybase.1.secretUi.getKeybasePassphrase", []interface{}{__arg}, &res)
	return
}

type Session struct {
	Uid      UID    `codec:"uid" json:"uid"`
	Username string `codec:"username" json:"username"`
	Token    string `codec:"token" json:"token"`
}

type CurrentSessionArg struct {
}

type SessionInterface interface {
	CurrentSession() (Session, error)
}

func SessionProtocol(i SessionInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.session",
		Methods: map[string]rpc2.ServeHook{
			"currentSession": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]CurrentSessionArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.CurrentSession()
				}
				return
			},
		},
	}

}

type SessionClient struct {
	Cli GenericClient
}

func (c SessionClient) CurrentSession() (res Session, err error) {
	err = c.Cli.Call("keybase.1.session.currentSession", []interface{}{CurrentSessionArg{}}, &res)
	return
}

type AddArg struct {
	SecretPhrase string `codec:"secretPhrase" json:"secretPhrase"`
}

type SibkeyInterface interface {
	Add(string) error
}

func SibkeyProtocol(i SibkeyInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.sibkey",
		Methods: map[string]rpc2.ServeHook{
			"add": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]AddArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Add(args[0].SecretPhrase)
				}
				return
			},
		},
	}

}

type SibkeyClient struct {
	Cli GenericClient
}

func (c SibkeyClient) Add(secretPhrase string) (err error) {
	__arg := AddArg{SecretPhrase: secretPhrase}
	err = c.Cli.Call("keybase.1.sibkey.add", []interface{}{__arg}, nil)
	return
}

type SignupRes struct {
	PassphraseOk bool `codec:"passphraseOk" json:"passphraseOk"`
	PostOk       bool `codec:"postOk" json:"postOk"`
	WriteOk      bool `codec:"writeOk" json:"writeOk"`
}

type CheckUsernameAvailableArg struct {
	Username string `codec:"username" json:"username"`
}

type SignupArg struct {
	Email      string `codec:"email" json:"email"`
	InviteCode string `codec:"inviteCode" json:"inviteCode"`
	Passphrase string `codec:"passphrase" json:"passphrase"`
	Username   string `codec:"username" json:"username"`
	DeviceName string `codec:"deviceName" json:"deviceName"`
}

type InviteRequestArg struct {
	Email    string `codec:"email" json:"email"`
	Fullname string `codec:"fullname" json:"fullname"`
	Notes    string `codec:"notes" json:"notes"`
}

type SignupInterface interface {
	CheckUsernameAvailable(string) error
	Signup(SignupArg) (SignupRes, error)
	InviteRequest(InviteRequestArg) error
}

func SignupProtocol(i SignupInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.signup",
		Methods: map[string]rpc2.ServeHook{
			"checkUsernameAvailable": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]CheckUsernameAvailableArg, 1)
				if err = nxt(&args); err == nil {
					err = i.CheckUsernameAvailable(args[0].Username)
				}
				return
			},
			"signup": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SignupArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.Signup(args[0])
				}
				return
			},
			"inviteRequest": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]InviteRequestArg, 1)
				if err = nxt(&args); err == nil {
					err = i.InviteRequest(args[0])
				}
				return
			},
		},
	}

}

type SignupClient struct {
	Cli GenericClient
}

func (c SignupClient) CheckUsernameAvailable(username string) (err error) {
	__arg := CheckUsernameAvailableArg{Username: username}
	err = c.Cli.Call("keybase.1.signup.checkUsernameAvailable", []interface{}{__arg}, nil)
	return
}

func (c SignupClient) Signup(__arg SignupArg) (res SignupRes, err error) {
	err = c.Cli.Call("keybase.1.signup.signup", []interface{}{__arg}, &res)
	return
}

func (c SignupClient) InviteRequest(__arg InviteRequestArg) (err error) {
	err = c.Cli.Call("keybase.1.signup.inviteRequest", []interface{}{__arg}, nil)
	return
}

type Sig struct {
	Seqno        int    `codec:"seqno" json:"seqno"`
	SigIdDisplay string `codec:"sigIdDisplay" json:"sigIdDisplay"`
	Type         string `codec:"type" json:"type"`
	Ctime        int    `codec:"ctime" json:"ctime"`
	Revoked      bool   `codec:"revoked" json:"revoked"`
	Active       bool   `codec:"active" json:"active"`
	Key          string `codec:"key" json:"key"`
	Body         string `codec:"body" json:"body"`
}

type SigTypes struct {
	Track          bool `codec:"track" json:"track"`
	Proof          bool `codec:"proof" json:"proof"`
	Cryptocurrency bool `codec:"cryptocurrency" json:"cryptocurrency"`
	Self           bool `codec:"self" json:"self"`
}

type SigListArgs struct {
	SessionID int       `codec:"sessionID" json:"sessionID"`
	Username  string    `codec:"username" json:"username"`
	AllKeys   bool      `codec:"allKeys" json:"allKeys"`
	Types     *SigTypes `codec:"types,omitempty" json:"types"`
	Filterx   string    `codec:"filterx" json:"filterx"`
	Verbose   bool      `codec:"verbose" json:"verbose"`
	Revoked   bool      `codec:"revoked" json:"revoked"`
}

type SigListArg struct {
	Arg SigListArgs `codec:"arg" json:"arg"`
}

type SigListJSONArg struct {
	Arg SigListArgs `codec:"arg" json:"arg"`
}

type SigsInterface interface {
	SigList(SigListArgs) ([]Sig, error)
	SigListJSON(SigListArgs) (string, error)
}

func SigsProtocol(i SigsInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.sigs",
		Methods: map[string]rpc2.ServeHook{
			"sigList": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SigListArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.SigList(args[0].Arg)
				}
				return
			},
			"sigListJSON": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]SigListJSONArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.SigListJSON(args[0].Arg)
				}
				return
			},
		},
	}

}

type SigsClient struct {
	Cli GenericClient
}

func (c SigsClient) SigList(arg SigListArgs) (res []Sig, err error) {
	__arg := SigListArg{Arg: arg}
	err = c.Cli.Call("keybase.1.sigs.sigList", []interface{}{__arg}, &res)
	return
}

func (c SigsClient) SigListJSON(arg SigListArgs) (res string, err error) {
	__arg := SigListJSONArg{Arg: arg}
	err = c.Cli.Call("keybase.1.sigs.sigListJSON", []interface{}{__arg}, &res)
	return
}

type CloseArg struct {
	S Stream `codec:"s" json:"s"`
}

type ReadArg struct {
	S  Stream `codec:"s" json:"s"`
	Sz int    `codec:"sz" json:"sz"`
}

type WriteArg struct {
	S   Stream `codec:"s" json:"s"`
	Buf []byte `codec:"buf" json:"buf"`
}

type StreamUiInterface interface {
	Close(Stream) error
	Read(ReadArg) ([]byte, error)
	Write(WriteArg) (int, error)
}

func StreamUiProtocol(i StreamUiInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.streamUi",
		Methods: map[string]rpc2.ServeHook{
			"close": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]CloseArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Close(args[0].S)
				}
				return
			},
			"read": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]ReadArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.Read(args[0])
				}
				return
			},
			"write": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]WriteArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.Write(args[0])
				}
				return
			},
		},
	}

}

type StreamUiClient struct {
	Cli GenericClient
}

func (c StreamUiClient) Close(s Stream) (err error) {
	__arg := CloseArg{S: s}
	err = c.Cli.Call("keybase.1.streamUi.close", []interface{}{__arg}, nil)
	return
}

func (c StreamUiClient) Read(__arg ReadArg) (res []byte, err error) {
	err = c.Cli.Call("keybase.1.streamUi.read", []interface{}{__arg}, &res)
	return
}

func (c StreamUiClient) Write(__arg WriteArg) (res int, err error) {
	err = c.Cli.Call("keybase.1.streamUi.write", []interface{}{__arg}, &res)
	return
}

type TrackArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	TheirName string `codec:"theirName" json:"theirName"`
}

type TrackInterface interface {
	Track(TrackArg) error
}

func TrackProtocol(i TrackInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.track",
		Methods: map[string]rpc2.ServeHook{
			"track": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]TrackArg, 1)
				if err = nxt(&args); err == nil {
					err = i.Track(args[0])
				}
				return
			},
		},
	}

}

type TrackClient struct {
	Cli GenericClient
}

func (c TrackClient) Track(__arg TrackArg) (err error) {
	err = c.Cli.Call("keybase.1.track.track", []interface{}{__arg}, nil)
	return
}

type PromptYesNoArg struct {
	SessionID int   `codec:"sessionID" json:"sessionID"`
	Text      Text  `codec:"text" json:"text"`
	Def       *bool `codec:"def,omitempty" json:"def"`
}

type UiInterface interface {
	PromptYesNo(PromptYesNoArg) (bool, error)
}

func UiProtocol(i UiInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.ui",
		Methods: map[string]rpc2.ServeHook{
			"promptYesNo": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]PromptYesNoArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.PromptYesNo(args[0])
				}
				return
			},
		},
	}

}

type UiClient struct {
	Cli GenericClient
}

func (c UiClient) PromptYesNo(__arg PromptYesNoArg) (res bool, err error) {
	err = c.Cli.Call("keybase.1.ui.promptYesNo", []interface{}{__arg}, &res)
	return
}

type Tracker struct {
	Tracker UID `codec:"tracker" json:"tracker"`
	Status  int `codec:"status" json:"status"`
	Mtime   int `codec:"mtime" json:"mtime"`
}

type TrackProof struct {
	ProofType string `codec:"proofType" json:"proofType"`
	ProofName string `codec:"proofName" json:"proofName"`
	IdString  string `codec:"idString" json:"idString"`
}

type TrackEntry struct {
	Username       string       `codec:"username" json:"username"`
	SigId          string       `codec:"sigId" json:"sigId"`
	PgpFingerprint string       `codec:"pgpFingerprint" json:"pgpFingerprint"`
	TrackTime      int64        `codec:"trackTime" json:"trackTime"`
	Proofs         []TrackProof `codec:"proofs" json:"proofs"`
}

type WebProof struct {
	Hostname  string   `codec:"hostname" json:"hostname"`
	Protocols []string `codec:"protocols" json:"protocols"`
}

type PubKey struct {
	KeyFingerprint string `codec:"keyFingerprint" json:"keyFingerprint"`
	Bits           int    `codec:"bits" json:"bits"`
	Algo           int    `codec:"algo" json:"algo"`
}

type Proofs struct {
	Social    []TrackProof `codec:"social" json:"social"`
	Web       []WebProof   `codec:"web" json:"web"`
	PublicKey PubKey       `codec:"publicKey" json:"publicKey"`
}

type UserSummary struct {
	Uid       UID    `codec:"uid" json:"uid"`
	Thumbnail string `codec:"thumbnail" json:"thumbnail"`
	Username  string `codec:"username" json:"username"`
	IdVersion int    `codec:"idVersion" json:"idVersion"`
	FullName  string `codec:"fullName" json:"fullName"`
	Bio       string `codec:"bio" json:"bio"`
	Proofs    Proofs `codec:"proofs" json:"proofs"`
}

type ListTrackersArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
	Uid       UID `codec:"uid" json:"uid"`
}

type ListTrackersByNameArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Username  string `codec:"username" json:"username"`
}

type ListTrackingArg struct {
	Filter string `codec:"filter" json:"filter"`
}

type ListTrackingJsonArg struct {
	Filter  string `codec:"filter" json:"filter"`
	Verbose bool   `codec:"verbose" json:"verbose"`
}

type LoadUncheckedUserSummariesArg struct {
	Uids []UID `codec:"uids" json:"uids"`
}

type UserInterface interface {
	ListTrackers(ListTrackersArg) ([]Tracker, error)
	ListTrackersByName(ListTrackersByNameArg) ([]Tracker, error)
	ListTracking(string) ([]TrackEntry, error)
	ListTrackingJson(ListTrackingJsonArg) (string, error)
	LoadUncheckedUserSummaries([]UID) ([]UserSummary, error)
}

func UserProtocol(i UserInterface) rpc2.Protocol {
	return rpc2.Protocol{
		Name: "keybase.1.user",
		Methods: map[string]rpc2.ServeHook{
			"listTrackers": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]ListTrackersArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.ListTrackers(args[0])
				}
				return
			},
			"listTrackersByName": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]ListTrackersByNameArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.ListTrackersByName(args[0])
				}
				return
			},
			"listTracking": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]ListTrackingArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.ListTracking(args[0].Filter)
				}
				return
			},
			"listTrackingJson": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]ListTrackingJsonArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.ListTrackingJson(args[0])
				}
				return
			},
			"loadUncheckedUserSummaries": func(nxt rpc2.DecodeNext) (ret interface{}, err error) {
				args := make([]LoadUncheckedUserSummariesArg, 1)
				if err = nxt(&args); err == nil {
					ret, err = i.LoadUncheckedUserSummaries(args[0].Uids)
				}
				return
			},
		},
	}

}

type UserClient struct {
	Cli GenericClient
}

func (c UserClient) ListTrackers(__arg ListTrackersArg) (res []Tracker, err error) {
	err = c.Cli.Call("keybase.1.user.listTrackers", []interface{}{__arg}, &res)
	return
}

func (c UserClient) ListTrackersByName(__arg ListTrackersByNameArg) (res []Tracker, err error) {
	err = c.Cli.Call("keybase.1.user.listTrackersByName", []interface{}{__arg}, &res)
	return
}

func (c UserClient) ListTracking(filter string) (res []TrackEntry, err error) {
	__arg := ListTrackingArg{Filter: filter}
	err = c.Cli.Call("keybase.1.user.listTracking", []interface{}{__arg}, &res)
	return
}

func (c UserClient) ListTrackingJson(__arg ListTrackingJsonArg) (res string, err error) {
	err = c.Cli.Call("keybase.1.user.listTrackingJson", []interface{}{__arg}, &res)
	return
}

func (c UserClient) LoadUncheckedUserSummaries(uids []UID) (res []UserSummary, err error) {
	__arg := LoadUncheckedUserSummariesArg{Uids: uids}
	err = c.Cli.Call("keybase.1.user.loadUncheckedUserSummaries", []interface{}{__arg}, &res)
	return
}
