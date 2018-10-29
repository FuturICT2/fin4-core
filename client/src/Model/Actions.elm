module Model.Actions exposing (Action, Actions, actionDecoder, actionsDecoder, init)

import Json.Decode as JD
import Json.Decode.Pipeline as JP


type alias Action =
    { name : String
    }


type alias Actions =
    { entries : List Action
    }


init : Actions
init =
    { entries = []
    }


actionsDecoder : JD.Decoder Actions
actionsDecoder =
    JP.decode Actions
        |> JP.required "entries" (JD.list actionDecoder)


actionDecoder : JD.Decoder Action
actionDecoder =
    JP.decode Action
        |> JP.required "Names" JD.string
