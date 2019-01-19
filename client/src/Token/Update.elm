module Token.Update exposing (update)

import Common.Json exposing (decodeAt, deocdeIntWithDefault)
import Common.Log exposing (log)
import Json.Decode as JD
import Main.Context exposing (Context)
import Material
import Token.Command
    exposing
        ( approveProposalCmd
        , loadTokenCmd
        , submitProposalCmd
        )
import Token.Model exposing (Model)
import Token.Msg exposing (Msg(..))
import Token.Ports exposing (fileSelected)


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        ToggleClaimButton ->
            { model | showClaimForm = not model.showClaimForm } ! []

        SelectTab tab ->
            { model | tab = tab, showClaimForm = False } ! []

        OnTokenLoadResponse resp ->
            case resp of
                Ok data ->
                    { model
                        | data = Just data
                    }
                        ! []

                Err _ ->
                    { model
                        | data = Nothing
                        , error = Just "Error loading token. Please try again."
                    }
                        ! []

        SubmitProposal proposal ->
            let
                tokenId =
                    case model.data of
                        Nothing ->
                            0

                        Just t ->
                            t.id
            in
            { model | proposalText = proposal } ! [ submitProposalCmd ctx model tokenId proposal ]

        SubmitProposalResponse resp ->
            case resp of
                Ok data ->
                    { model
                        | proposalText = ""
                        , showClaimForm = False
                        , submitProposalError = Nothing
                    }
                        ! []

                Err error ->
                    { model
                        | submitProposalError = Just error
                    }
                        ! []

        ApproveClaim claimId ->
            model ! [ approveProposalCmd ctx model claimId ]

        ApproveClaimResponse resp ->
            model ! []

        SetClaim proposalText ->
            { model | proposalText = proposalText } ! []

        ImageSelected ->
            model ! [ fileSelected <| "claim-img-proposal" ]

        ImageRead data ->
            let
                newImage =
                    { contents = data.contents
                    , filename = data.filename
                    }
            in
            { model
                | img = Just newImage
            }
                ! []

        OnDoLikeResponse resp ->
            model ! []

        DoLike state ->
            model ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
