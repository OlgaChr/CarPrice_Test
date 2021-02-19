<?php

use App\Http\Controllers\AuthorController;
use App\Http\Controllers\BookController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/


Route::get('authors', [AuthorController::class, 'index']);
Route::get('author/{author}', [AuthorController::class, 'show']);
Route::post('author', [AuthorController::class, 'store']);
Route::put('author/{author}', [AuthorController::class, 'update']);
Route::delete('author/{author}', [AuthorController::class, 'destroy']);

Route::get('books', [BookController::class, 'index']);
Route::get('book/{book}', [BookController::class, 'show']);
Route::post('book', [BookController::class, 'store']);
Route::put('book/{book}', [BookController::class, 'update']);
Route::delete('book/{book}', [BookController::class, 'destroy']);
