module Timeline.Msg exposing (Msg(..))

import Common.Json exposing (EmptyResponse)
import Gallery
import Http
import Material
import Timeline.Model exposing (TimelineEntries)


type Msg
    = Mdl (Material.Msg Msg)
    | OnLoadTimelineResponse (Result Http.Error TimelineEntries)
    | GalleryMsg Int Gallery.Msg
    | LoadMore
    | OnToggleFavoriteResponse (Result Http.Error {})
    | ToggleFavorite Int
    | VerifyBlock Bool Int
    | VerifyBlockResponse (Result Http.Error EmptyResponse)
