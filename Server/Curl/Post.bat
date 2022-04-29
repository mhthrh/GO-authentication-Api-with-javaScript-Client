
@REM :start
@REM :: Post Add Score
@REM 	curl -Lvso /dev/null -d  "@sign.json" -X POST http://localhost:8585/scores
@REM 	curl GET http://localhost:8585/signup/49385234
@REM goto start

@REM curl -Lvso /dev/null -d  "@SignUp.json" -X POST http://localhost:8585/signUp
 curl -Lvso /dev/null -d  "@SignIn.json" -X POST http://localhost:8585/signIn