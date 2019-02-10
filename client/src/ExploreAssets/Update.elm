module ExploreAssets.Update exposing (update)

import ExploreAssets.Command exposing (loadFavoritesCmd, toggleFavoriteCmd)
import ExploreAssets.Model exposing (Model)
import ExploreAssets.Msg exposing (Msg(..))
import Main.Context exposing (Context)
import Material


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        OnLoadAssetsResponse resp ->
            case resp of
                Ok assets ->
                    { model | assets = assets, error = Nothing } ! []

                Err _ ->
                    { model | error = Just "Error loading assets" } ! []

        OnLoadFavoritesResponse resp ->
            case resp of
                Ok favs ->
                    { model | favorites = favs, error = Nothing } ! []

                Err _ ->
                    { model | error = Just "Error loading favorites" } ! []

        ToggleFavorite id ->
            model ! [ toggleFavoriteCmd ctx id ]

        OnToggleFavoriteResponse resp ->
            case resp of
                Ok _ ->
                    model ! [ loadFavoritesCmd ctx ]

                Err _ ->
                    model ! []

        OnSwitchCategory category ->
            { model | activeCategory = category } ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
