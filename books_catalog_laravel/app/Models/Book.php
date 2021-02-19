<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Book extends Model
{
    use HasFactory;

    protected $fillable = ['name', 'publication_year', 'summary'];
    protected $hidden = ['created_at', 'updated_at', 'deleted_at'];
    protected $appends = ['authors'];

    public function authors()
    {
        return $this->belongsToMany(
            Author::class,
            'authors_books',
            'book_id',
            'author_id');
    }

    public function getAuthorsAttribute()
    {
        return $this->authors()->get();
    }
}
