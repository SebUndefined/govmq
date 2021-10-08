package hook

// AuthOnRegister : The auth_on_register hook allow your plugin to grant or reject new client connections.
//Moreover, it lets you exert fine-grained control over the configuration of the client session.
//The auth_on_register hook is specified in the Erlang behaviour auth_on_register_hook
//(https://github.com/vernemq/vernemq_dev/blob/master/src/auth_on_register_hook.erl)
// available in the vernemq_dev repo.
//Every plugin that implements the auth_on_register hook are part of a conditional plugin chain.
//For this reason we allow the hook to return different values depending on how the plugin grants
//or rejects this client. In case the plugin doesn't know the client it is best to return next as
//this would allow subsequent plugins in the chain to validate this client. If no plugin is able
//to validate the client it gets automatically rejected.
type AuthOnRegister struct {
	MountPoint   string `json:"mountpoint",validate:"required"`
	ClientId     string `json:"client_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PeerAddress  string `json:"peer_addr"`
	PeerPort     int    `json:"peer_port"`
	CleanSession bool   `json:"clean_session"`
}

// OnClientWakeUp : Once a new client was successfully authenticated and the above
//described hooks have been called, the client attaches to its queue. If it is a returning client
//using clean_session=false or if the client had previous sessions in the cluster,
//this process could take a while. (As offline messages are migrated to a new node, existing
//sessions are disconnected).
//The on_client_wakeup (https://github.com/vernemq/vernemq_dev/blob/master/src/on_client_wakeup_hook.erl)
//hook is called at the point where a queue has been successfully instantiated, possible offline messages
//migrated, and potential duplicate sessions have been disconnected. In other words: when
//the client has reached a completely initialized, normal state for accepting messages.
//The hook is specified in the Erlang behaviour on_client_wakeup_hook available in the vernemq_dev repo.
type OnClientWakeUp struct {
	MountPoint string `json:"mountpoint"`
	ClientId   string `json:"client_id"`
}

// OnRegister : The on_register hook allow your plugin to get informed about
//a newly authenticated client. The hook is specified in the Erlang behaviour on_register_hook
// (https://github.com/vernemq/vernemq_dev/blob/master/src/on_register_hook.erl)
//available in the vernemq_dev repo.
type OnRegister struct {
	MountPoint  string `json:"mountpoint"`
	ClientId    string `json:"client_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PeerAddress string `json:"peer_address"`
}

// OnClientOffline : This hook is called if an MQTT 3.1/3.1.1 client using clean_session=false
//or an MQTT 5.0 client with a non-zero session_expiry_interval closes the connection or gets
//disconnected by a duplicate client. The hook is specified in the Erlang behaviour on_client_offline_hook
//(https://github.com/vernemq/vernemq_dev/blob/master/src/on_client_offline_hook.erl)
//available in the vernemq_dev repo.
type OnClientOffline struct {
	MountPoint string `json:"mountpoint"`
	ClientId   string `json:"client_id"`
}

// OnClientGone : This hook is called if an MQTT 3.1/3.1.1 client using clean_session=true or an MQTT 5.0
//client with the session_expiry_interval set to zero closes the connection or gets disconnected by a
//duplicate client. The hook is specified in the Erlang behaviour on_client_gone_hook
//(https://github.com/vernemq/vernemq_dev/blob/master/src/on_client_gone_hook.erl)
//available in the vernemq_dev repo.
type OnClientGone struct {
	MountPoint string `json:"mountpoint"`
	ClientId   string `json:"client_id"`
}
