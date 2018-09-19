module Common.Task exposing (delay)

import Time
import Task
import Process


{- use this to delay a task in update(for example) -}


delay : Time.Time -> msg -> Cmd msg
delay time msg =
    Process.sleep time
        |> Task.perform (\_ -> msg)
