package domain

type ProviderHandle struct{ Provider PolicyProvider }

func (h ProviderHandle) Available() bool { return h.Provider != nil }
