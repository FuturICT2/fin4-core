module CreateAction.Update exposing (update)

import CreateAction.Command exposing (createActionCmd)
import CreateAction.Model exposing (Model)
import CreateAction.Msg exposing (Msg(..))
import Debug
import List.Extra
import Main.Context exposing (Context)
import Material
import Navigation exposing (newUrl)


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        PostAction ->
            { model
                | isCreatingAction = True
                , createActionError = Nothing
            }
                ! [ createActionCmd ctx model ]

        OnCreateActionResponse res ->
            case res of
                Ok _ ->
                    { model
                        | isCreatingAction = False
                        , createActionError = Nothing
                        , step = model.step + 1
                    }
                        ! [ newUrl "#tokens" ]

                Err error ->
                    Debug.log (toString error)
                        { model
                            | isCreatingAction = False
                            , createActionError = Just error
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

        SetTimeLimit value ->
            { model | timeLimit = value } ! []

        SetDescription value ->
            { model | description = value } ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
