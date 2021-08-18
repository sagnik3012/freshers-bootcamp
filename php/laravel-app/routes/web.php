<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\StudentController;
/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return view('welcome');
});



Route::post('student',       [StudentController::class, 'createStudent']);
Route::get('student/{id}',   [StudentController::class, 'getStudentById']);
Route::get('student',        [StudentController::class, 'getAllStudents']);
Route::delete('student/{id}',[StudentController::class, 'deleteStudentById']);
