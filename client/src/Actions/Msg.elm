module Actions.Msg exposing (Msg(..))

import Common.Json exposing (EmptyResponse)
import Http
import Material
import Model.Actions exposing (Actions)


type Msg
    = Mdl (Material.Msg Msg)
    | OnLoadActionsResponse (Result Http.Error Actions)
    | TickerTimout
    | OnDoLikeResponse (Result Http.Error EmptyResponse)
    | DoLike Int Bool
