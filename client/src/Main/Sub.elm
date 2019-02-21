module Main.Sub exposing (subscriptions)

import Asset.Sub
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing exposing (Route(..))
import Profile.Sub
import Window exposing (..)


subscriptions : Model -> Sub Msg
subscriptions model =
    let
        sub =
            case model.context.route of
                AssetRoute assetId ->
                    Sub.batch
                        [ Sub.map AssetMsg <|
                            Asset.Sub.subscriptions model.context
                        ]

                ProfileRoute profileId ->
                    Sub.batch
                        [ Sub.map ProfileMsg <|
                            Profile.Sub.subscriptions model.context
                        ]

                _ ->
                    Sub.none
    in
    Sub.batch
        [ Window.resizes (\size -> OnWindowResize size)
        , sub
        ]
