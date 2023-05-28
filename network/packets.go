package network

type askInfo struct {
  Version uint8
  Time uint32
}

type ServerInfo struct {
  OldVersion uint8
  PacketVersion uint8
  Application [16]byte
  Version, Subversion, NumberOfPlayer, MaxPlayer, Gametype, ModifiedGame, CheatsEnabled, KartVars, FileNeededNum uint8
  Time, LevelTime uint32
  ServerName [32]byte
  MapName [8]byte
  MapTitle [33]byte
  MapMd5 [16]byte
  ActNum uint8
  IsZone uint8
  HttpSource [256]byte
  FileNeeded [915]uint8
}

type PlayerInfo struct {
  Node uint8
  Name [22]byte
  Address [4]uint8
  Team, Skin, Data uint8
  Score uint32
  TimeInServer uint16
}
