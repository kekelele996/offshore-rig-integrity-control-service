package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
func(s *System) StageManifest(id string)error{zones,err:=s.Registry.DispatchSnapshot(id);if err!=nil{return err};s.Manifests.Save(id,zones);return nil}
func(s *System) AddEmergencyZone(id,zone string)[]string{zones:=s.Manifests.Read(id);zones=append(zones,zone);return domain.CloneStrings(zones)}
func(s *System) StoredManifest(id string)[]string{return s.Manifests.Read(id)}
