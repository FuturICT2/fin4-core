module Asset.Msg exposing (Msg(..))

import Asset.Model exposing (Asset)
import Asset.Ports exposing (ImagePortData)
import Common.Json exposing (EmptyResponse)
import Http
import Material
import Timeline.Msg


type Msg
    = Mdl (Material.Msg Msg)
    | OnAssetLoadResponse (Result Http.Error Asset)
    | SelectTab Int
    | SetBlockText String
    | SubmitBlock
    | SubmitBlockResponse (Result Http.Error EmptyResponse)
    | VerifyBlock Bool Int
    | VerifyBlockResponse (Result Http.Error EmptyResponse)
    | ImageSelected
    | ImageRead ImagePortData
    | DeleteImage Int
    | TimelineMsg Timeline.Msg.Msg
