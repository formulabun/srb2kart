package network

type packettype_t uint8
// from d_clisrv.h
const (
	pt_nothing   packettype_t = iota // To send a nop through the network. ^_~
	pt_servercfg              // Server config used in start game
	// (must stay 1 for backwards compatibility).
	// This is a positive response to a CLIENTJOIN request.
	pt_clientcmd     // Ticcmd of the client.
	pt_clientmis     // Same as above with but saying resend from.
	pt_client2cmd    // 2 cmds in the packet for splitscreen.
	pt_client2mis    // Same as above with but saying resend from
	pt_nodekeepalive // Same but without ticcmd and consistancy
	pt_nodekeepalivemis
	pt_servertics   // All cmds for the tic.
	pt_serverrefuse // Server refuses joiner (reason inside).
	pt_servershutdown
	pt_clientquit // Client closes the connection.

	pt_askinfo      // Anyone can ask info of the server.
	pt_serverinfo   // Send game & server info (gamespy).
	pt_playerinfo   // Send information for players in game (gamespy).
	pt_requestfile  // Client requests a file transfer
	pt_askinfoviams // Packet from the MS requesting info be sent to new client.
	// If this ID changes update masterserver definition.
	pt_resynchend // Player is now resynched and is being requested to remake the gametic
	pt_resynchget // Player got resynch packet

	// Add non-PT_CANFAIL packet types here to avoid breaking MS compatibility.

	// Kart-specific packets
	pt_client3cmd // 3P
	pt_client3mis
	pt_client4cmd // 4P
	pt_client4mis
	pt_basickeepalive // Keep the network alive during wipes, as tics aren't advanced and NetUpdate isn't called

	pt_canfail // This is kind of a priority. Anything bigger than CANFAIL
	// allows HSendPacket(* true, *, *) to return false.
	// In addition this packet can't occupy all the available slots.


	pt_textcmd     // Extra text commands from the client.
	pt_textcmd2    // Splitscreen text commands.
	pt_textcmd3    // 3P
	pt_textcmd4    // 4P
	pt_clientjoin  // Client wants to join; used in start game.
	pt_nodetimeout // Packet sent to self if the connection times out.
	pt_resynching  // Packet sent to resync players.
	// Blocks game advance until synched.

	pt_tellfilesneeded // Client, to server: "what other files do I need starting from this number?"
	pt_morefilesneeded // Server, to client: "you need these (+ more on top of those)"

	pt_ping // packet sent to tell clients the other client's latency to server.
	numpackettype
)
const pt_filefragment = pt_canfail // A part of a file.
