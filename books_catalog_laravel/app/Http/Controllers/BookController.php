<?php

namespace App\Http\Controllers;

use App\Http\Requests\BookRequest;
use App\Models\Book;
use Illuminate\Http\Request;
use Illuminate\Validation\ValidationException;

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
     * @param \Illuminate\Http\Request $request
     * @return \Illuminate\Http\Response
     * @throws ValidationException
     */
    public function store(BookRequest $request)
    {
        $authors = $request->get('authors');
        if (!$authors || !is_array($authors)) {
            throw ValidationException::withMessages(['authors' => 'This value required, type = array of ids']);
        }
        $book = Book::create($request->except(['authors']));
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
    public function update(BookRequest $request, Book $book)
    {
        $authors = $request->get('authors');
        $book->update($request->except(['authors']));
        $book->authors()->sync($authors);
        return $book;
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param \App\Models\Book $book
     * @return \Illuminate\Http\Response
     * @throws \Exception
     */
    public function destroy(Book $book)
    {
        $book->authors()->detach();
        if ($book->delete()) {
            return response(null, 204);
        }
    }
}
