module Main.Sub exposing (subscriptions)

import Actions.Sub
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing exposing (Route(..))
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

                ActionsRoute ->
                    Sub.batch
                        [ Sub.map Actions <|
                            Actions.Sub.subscriptions model.context
                        ]

                _ ->
                    Sub.none
    in
    Sub.batch
        [ Window.resizes (\size -> OnWindowResize size)
        , sub
        ]
