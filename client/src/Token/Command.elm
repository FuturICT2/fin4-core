module Token.Command exposing
    ( approveProposalCmd
    , commands
    , likeCmd
    , loadTokenCmd
    , submitProposalCmd
    )

import Common.Api exposing (deleteWithCsrf, get, postWithCsrf)
import Common.Json exposing (decodeAt, deocdeIntWithDefault, emptyResponseDecoder)
import Json.Decode as JD
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Model.Tokens exposing (tokenDecoder)
import Token.Model exposing (Model)
import Token.Msg exposing (Msg(..))


commands : Context -> Int -> Cmd Msg
commands ctx tokenId =
    case ctx.route of
        TokenRoute tokenId ->
            loadTokenCmd ctx tokenId

        _ ->
            Cmd.none


loadTokenCmd ctx tokenId =
    get ctx
        OnTokenLoadResponse
        ("/tokens/" ++ toString tokenId)
        []
        tokenDecoder


approveProposalCmd : Context -> Model -> Int -> Cmd Msg
approveProposalCmd ctx model claimId =
    postWithCsrf ctx
        ApproveClaimResponse
        "/approve-claim"
        (encodeApproveProposal claimId)
        emptyResponseDecoder


submitProposalCmd : Context -> Model -> Int -> String -> Cmd Msg
submitProposalCmd ctx model actionId proposal =
    let
        v =
            Maybe.withDefault { contents = "", filename = "" } model.img
    in
    postWithCsrf ctx
        SubmitProposalResponse
        "/create-claim"
        (encodeSubmitProposal actionId proposal v.contents)
        emptyResponseDecoder


likeCmd ctx tokenId state =
    let
        url =
            case state of
                True ->
                    "/unlike/" ++ toString tokenId

                False ->
                    "/like/" ++ toString tokenId
    in
    get ctx
        OnDoLikeResponse
        url
        []
        emptyResponseDecoder


encodeApproveProposal : Int -> JE.Value
encodeApproveProposal claimId =
    JE.object
        [ ( "claimId", JE.int claimId )
        ]


encodeSubmitProposal tokenId claim base64 =
    JE.object
        [ ( "tokenId", JE.int tokenId )
        , ( "proposal", JE.string claim )
        , ( "image64", JE.string base64 )
        ]
