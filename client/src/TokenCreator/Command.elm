module TokenCreator.Command exposing (commands)

import Common.Api exposing (get)
import Common.Json exposing (decodeAt, deocdeIntWithDefault)
import Json.Decode as JD
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import TokenCreator.Msg exposing (Msg(..))


commands : Context -> Cmd Msg
commands ctx =
    case ctx.route of
        TokenCreatorRoute ->
            Cmd.none

        _ ->
            Cmd.none
