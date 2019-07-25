//////////////////////////////////////////////////////////////////////////////
//
// Config contains configuration data for PeerConnection
//
// Copyright 2019 Lanikai Labs. All rights reserved.
//
//////////////////////////////////////////////////////////////////////////////

package alohartc

import (
	"github.com/lanikai/alohartc/internal/media"
)

type Config struct {
	LocalVideo media.VideoSource
	//AudioTrack Track
	//VideoTrack Track
}
