module CreateAsset.Msg exposing (Msg(..))

import Asset.Model exposing (Asset)
import CreateAsset.Types exposing (..)
import Http
import Material


type Msg
    = Mdl (Material.Msg Msg)
    | PostAsset
    | OnCreateAssetSuccess (Result Http.Error Asset)
    | SetName String
    | SetSymbol String
    | SetDescription String
    | SetActiveView ActiveView
