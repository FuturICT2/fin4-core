module ExploreAssets.Command exposing (commands, loadAssetsCmd, loadFavoritesCmd, toggleFavoriteCmd)

import Asset.Model exposing (assetListDecoder)
import Common.Api exposing (get, postWithCsrf)
import Common.Json exposing (decodeAt, deocdeIntWithDefault, emptyResponseDecoder)
import ExploreAssets.Msg exposing (Msg(..))
import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))


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


loadAssetsCmd ctx page =
    get ctx
        OnLoadAssetsResponse
        "/assets"
        [ ( "page", toString page ) ]
        assetListDecoder


loadFavoritesCmd ctx =
    get ctx
        OnLoadFavoritesResponse
        "/assets-favorites"
        []
        assetListDecoder


toggleFavoriteCmd ctx assetId =
    postWithCsrf ctx
        OnToggleFavoriteResponse
        ("/assets/" ++ toString assetId ++ "/toggle-favorite")
        (JE.object
            [ ( "assetId", JE.int assetId )
            ]
        )
        (JP.decode {})
