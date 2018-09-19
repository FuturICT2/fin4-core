module Main.Sub exposing (subscriptions)

import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing exposing (Route(..))
import WebSocket
import Window exposing (..)


subscriptions : Model -> Sub Msg
subscriptions model =
    let
        sub =
            Sub.none
    in
    Sub.batch
        [ Window.resizes (\size -> OnWindowResize size)
        , sub
        ]
