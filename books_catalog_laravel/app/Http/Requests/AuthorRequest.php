<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class AuthorRequest extends FormRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'surname' => 'required|string',
            'name' => 'required|string',
            'middlename' => 'string',
            'birth_year' => 'required|integer',
            'death_year' => 'integer|lt:birth_year',
        ];
    }
}
