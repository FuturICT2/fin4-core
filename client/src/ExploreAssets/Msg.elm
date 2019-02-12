module ExploreAssets.Msg exposing (Msg(..))

import Asset.Model exposing (Asset)
import Http
import Material


type Msg
    = Mdl (Material.Msg Msg)
    | OnLoadAssetsResponse (Result Http.Error (List Asset))
    | OnLoadFavoritesResponse (Result Http.Error (List Asset))
    | OnToggleFavoriteResponse (Result Http.Error {})
    | ToggleFavorite Int
    | OnSwitchCategory String
