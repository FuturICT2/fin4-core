module CreateToken.Update exposing (update)

import CreateToken.Command exposing (createTokenCmd)
import CreateToken.Model exposing (Model)
import CreateToken.Msg exposing (Msg(..))
import Debug
import List.Extra
import Main.Context exposing (Context)
import Material


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        PostToken ->
            { model
                | isCreatingToken = True
                , createTokenError = Nothing
            }
                ! [ createTokenCmd ctx model ]

        OnCreateTokenResponse res ->
            case res of
                Ok token ->
                    { model
                        | isCreatingToken = False
                        , createTokenError = Nothing
                        , step = model.step + 1
                    }
                        ! []

                Err error ->
                    Debug.log (toString error)
                        { model
                            | isCreatingToken = False
                            , createTokenError = Just error
                        }
                        ! []

        StepBack ->
            let
                step =
                    case model.step - 1 >= 0 of
                        True ->
                            model.step - 1

                        False ->
                            0
            in
            { model | step = step } ! []

        StepForward ->
            { model | step = model.step + 1 } ! []

        SetName value ->
            { model | name = value } ! []

        SetSymbol value ->
            { model | symbol = value } ! []

        SetDescription value ->
            { model | description = value } ! []

        SetShares value ->
            { model | shares = value } ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
