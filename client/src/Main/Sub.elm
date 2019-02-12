module Main.Sub exposing (subscriptions)

import Asset.Sub
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing exposing (Route(..))
import Profile.Sub
import Token.Sub
import Tokens.Sub
import WebSocket
import Window exposing (..)


subscriptions : Model -> Sub Msg
subscriptions model =
    let
        sub =
            case model.context.route of
                TokensRoute ->
                    Sub.batch
                        [ Sub.map Tokens <|
                            Tokens.Sub.subscriptions model.context
                        ]

                TokenRoute tokenId ->
                    Sub.batch
                        [ Sub.map Token <|
                            Token.Sub.subscriptions model.context
                        ]

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
