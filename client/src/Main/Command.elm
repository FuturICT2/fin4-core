module Main.Command exposing (checkSessionCmd, commands)

import Common.Api exposing (get, postWithCsrf)
import Common.Json exposing (emptyResponseDecoder)
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.User exposing (User, userDecoder)
import Task exposing (Task)
import Window


commands : Context -> Cmd Msg
commands ctx =
    Cmd.batch
        [ -- subscribes to window size changes
          Task.perform (\size -> OnWindowResize size) Window.size
        , checkSessionCmd ctx
        ]


checkSessionCmd : Context -> Cmd Msg
checkSessionCmd ctx =
    get ctx OnCheckSessionResponse "/session" [] userDecoder
