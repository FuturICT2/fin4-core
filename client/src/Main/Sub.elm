module Main.Sub exposing (subscriptions)

import Tokens.Sub
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing exposing (Route(..))
import Token.Sub
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

                _ ->
                    Sub.none
    in
    Sub.batch
        [ Window.resizes (\size -> OnWindowResize size)
        , sub
        ]
