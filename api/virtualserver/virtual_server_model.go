package virtualserver

// VirtualServer data structure
type VirtualServer struct {
	Properties Properties `json:"properties"`
}

// Properties : Properties contains the overall monitor configuration
type Properties struct {
	Basic Basic `json:"basic"`
}

// Basic : Basic virtual server configration
type Basic struct {
	Enabled             bool        `json:"enabled"`
	DefaultTrafficPool  string      `json:"pool"`
	Port                int         `json:"port"`
	Protocol            string      `json:"protocol"`
	// all following attributes have default values set at
    // creation time if no value is specified...
    AddClusterIp        bool        `json:"add_cluster_ip"`         // true
    AddXForwardedFor    bool        `json:"add_x_forwarded_for"`    // false
    AddXForwardedProto  bool        `json:"add_x_forwarded_proto"`  // false
    BandwidthClass      string      `json:"bandwidth_class"`        // ""
    CloseWithRst        bool        `json:"close_with_rst"`         // false
    CompletionRules     []string    `json:"completionrules"`        // []
    ConnectTimeout      int         `json:"connect_timeout"`        // 10
	FtpForceServerSecure bool       `json:"ftp_force_server_secure"` // true

	// all these attributes are not mandatory at creation time...
	/*
	   ListenOnTrafficIps      []string        `json:"listen_on_traffic_ips"`
	   Note                    string          `json:"note"`
	   RequestRules            []string        `json:"request_rules"`
	   ResponseRules           []string        `"response_rules"`
	   FtpForceServerSecure    bool            `json:"ftp_force_server_secure"`
	   // SslClientCertHeaders : enum can accept these strings("all", "none"
	   // (default), "simple"
	   SslDecrypt              bool            `json:"ssl_decrypt"`
	   SslClientCertHeaders    string          `json:"ssl_client_cert_headers"`
	   SlmClass                string          `json:"slm_class"`
	   Transparent             bool            `json:"transparent"`
	   AddxforwardedProto      bool            `json:"add_x_forwarded_proto"`
	   CloseWithRst            bool            `json:"close_with_rst"`
	   // TBC...
	*/
}

// VirtualServersList : List of nodes monitored
type VirtualServersList struct {
	Children []ChildVirtualServer `json:"children"`
}

// ChildVirtualServer : monitored node structure
type ChildVirtualServer struct {
	Name string `json:"name"`
	HRef string `json:"href"`
}
