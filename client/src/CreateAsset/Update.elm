module CreateAsset.Update exposing (update)

import CreateAsset.Command exposing (createAssetCmd)
import CreateAsset.Model exposing (Model)
import CreateAsset.Msg exposing (Msg(..))
import CreateAsset.Types exposing (..)
import Main.Context exposing (Context)
import Material


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        SetActiveView view ->
            { model | step = view } ! []

        PostAsset ->
            { model
                | isCreatingAsset = True
                , createAssetError = Nothing
            }
                ! [ createAssetCmd ctx model ]

        OnCreateAssetSuccess res ->
            case res of
                Ok asset ->
                    { model
                        | isCreatingAsset = False
                        , createAssetError = Nothing
                        , step = SuccessView
                        , createdAssetName = asset.name
                        , createdAssetSymbol = asset.symbol
                        , createdAssetId = asset.id
                    }
                        ! []

                Err error ->
                    { model
                        | isCreatingAsset = False
                        , createAssetError = Just error
                    }
                        ! []

        SetName value ->
            { model | name = value } ! []

        SetSymbol value ->
            { model | symbol = value } ! []

        SetDescription value ->
            { model | purpose = value } ! []

        ToggleOracleType ->
            { model | isSensor = not model.isSensor } ! []
        
        ToggleIsBurnable ->
            { model | isBurnable = not model.isBurnable } ! []
        
        ToggleIsTransferable ->
            { model | isTransferable = not model.isTransferable } ! []
        
        ToggleIsMintable ->
            { model | isMintable = not model.isMintable } ! []
        
        Mdl msg_ ->
            Material.update Mdl msg_ model
