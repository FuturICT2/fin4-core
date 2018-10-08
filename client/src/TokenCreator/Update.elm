module TokenCreator.Update exposing (update)

import Debug
import Main.Context exposing (Context)
import Material


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        Mdl msg_ ->
            Material.update Mdl msg_ model

        SetName value ->
            { model | name = value } ! []

        SetTotalSupply value ->
            { model | totalSupply = value } ! []
