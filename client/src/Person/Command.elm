module Person.Command exposing
    ( commands
    , loadPersonCmd
    )

import Common.Api exposing (deleteWithCsrf, get, postWithCsrf)
import Common.Json exposing (decodeAt, deocdeIntWithDefault, emptyResponseDecoder)
import Json.Decode as JD
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Model.Persons exposing (personDecoder)
import Person.Model exposing (Model)
import Person.Msg exposing (Msg(..))


commands : Context -> Int -> Cmd Msg
commands ctx tokenId =
    case ctx.route of
        PersonRoute personId ->
            loadPersonCmd ctx personId

        _ ->
            Cmd.none


loadPersonCmd : Context -> Int -> Cmd Msg
loadPersonCmd ctx personId =
    get ctx
        OnPersonLoadResponse
        ("/person/" ++ toString personId)
        []
        personDecoder
