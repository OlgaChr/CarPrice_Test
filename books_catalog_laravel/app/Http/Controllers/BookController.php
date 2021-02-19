<?php

namespace App\Http\Controllers;

use App\Models\Book;
use Illuminate\Http\Request;

class BookController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @param Request $request
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        $author_id = $request->query('author_id');
        $size = $request->query('size');
        $query = Book::query();

        if ($author_id) {
            $query = $query->whereHas('authors', function ($q) use ($author_id){
                $q->where('authors.id', $author_id);
            });
        }

        if ($size) {
            return $query->paginate($size);
        }

        return $query->get();
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        $authors = $request->get('authors');
        $book = Book::create($request->except(['authors'])/*->validated()*/);
        $book->authors()->attach($authors);
        return $book;
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Book  $book
     * @return \Illuminate\Http\Response
     */
    public function show(Book $book)
    {
        return $book;
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\Book  $book
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Book $book)
    {
        $authors = $request->get('authors');
        $book->update($request->except(['authors']));
        $book->authors()->sync($authors);
        return $book;
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Book  $book
     * @return \Illuminate\Http\Response
     */
    public function destroy(Book $book)
    {
        $book->authors()->detach();
        if ($book->delete()) {
            return response(null, 204);
        }
    }
}
