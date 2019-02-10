module ExploreAssets.Command exposing
    ( commands
    , loadAssetsCmd
    , loadFavoritesCmd
    , toggleFavoriteCmd
    )

import Common.Api exposing (get, postWithCsrf)
import ExploreAssets.Msg exposing (Msg(..))
import Json.Decode.Pipeline as JP
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Model.Asset exposing (assetListDecoder)


commands : Context -> Cmd Msg
commands ctx =
    case ctx.route of
        ExploreAssetsRoute ->
            Cmd.batch
                [ loadAssetsCmd ctx 1
                , loadFavoritesCmd ctx
                ]

        _ ->
            Cmd.none


loadAssetsCmd : Context -> Int -> Cmd Msg
loadAssetsCmd ctx page =
    get ctx
        OnLoadAssetsResponse
        "/assets"
        [ ( "page", toString page ) ]
        assetListDecoder


loadFavoritesCmd : Context -> Cmd Msg
loadFavoritesCmd ctx =
    get ctx
        OnLoadFavoritesResponse
        "/assets-favorites"
        []
        assetListDecoder


toggleFavoriteCmd : Context -> Int -> Cmd Msg
toggleFavoriteCmd ctx assetId =
    postWithCsrf ctx
        OnToggleFavoriteResponse
        "/assets-favorites/toggle"
        (JE.object
            [ ( "assetId", JE.int assetId )
            ]
        )
        (JP.decode {})
