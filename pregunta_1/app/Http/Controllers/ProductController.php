<?php

namespace App\Http\Controllers;

use App\Models\Product;
use Illuminate\Http\Request;
use Validator;

class ProductController extends Controller
{
    public function getAll()
    {
        // obtener la lista todal de los productos
        return response()->json(Product::all(), 200);
    }
    public function getById(Request $request)
    {
        // obtener un producto segun la id
        $product = Product::find($request->id);
        if (!$product) {
            // en caso el producto no exista devuelve el siguiente response
            $response = [
                'id' => $request->id,
                'message' => 'Producto no existente'
            ];
            return response()->json($response, 404);
        }
        return response()->json($product, 200);
    }
    public function store(Request $request)
    {
        // reglas de validacion para el json de envio
        $validated = $request->validate([
            'name' => 'required|string|max:255',
            'description' => 'required|string',
            'price' => 'required|numeric|min:0',
            'stock' => 'required|integer|min:0',
        ]);
        // en caso de no se respete ninguna de las reglas, se crea la excepcion
        if (!$validated) {
            $response = [
                'id' => $request->id,
                'message' => 'No se esta respetando el formato preestablecido',
            ];
            return response()->json($response, 404);
        }

        $product = new Product();
        $product->name = $validated['name'];
        $product->description = $validated['description'];
        $product->price = $validated['price'];
        $product->stock = $validated['stock'];
        //guarda el producto
        $product->save();

        return response()->json($product, 201);
    }
    public function update(Request $request)
    {
        // reglas de validacion para el json de envio
        $validated = $request->validate([
            'name' => 'required|string|max:255',
            'description' => 'required|string',
            'price' => 'required|numeric|min:0',
            'stock' => 'required|integer|min:0',
        ]);
        // obtiene el producto segun el id
        $existsProduct = Product::where('id', $request->id)->first();
        // valida en caso exista el producto
        if (!$existsProduct) {
            $response = [
                'id' => $request->id,
                'message' => 'Producto no existente',
            ];
            return response()->json($response, 404);
        }

        $existsProduct->name = $validated['name'];
        $existsProduct->description = $validated['description'];
        $existsProduct->price = $validated['price'];
        $existsProduct->stock = $validated['stock'];
        // guarda el producto actualizado
        $existsProduct->save();
        $response = [
            'id' => $request->id,
            'message' => 'Producto actualizado',
            'data' => $validated
        ];
        return response()->json($response, 201);
    }
    public function delete(Request $request)
    {
        $product = Product::find($request->id);
        // valida si el producto existe o no
        if (!$product) {
            $response = [
                'id' => $request->id,
                'message' => 'Producto no existente'
            ];
            return response()->json($response, 400);
        }
        // eliminar totalmente de la base de datos
        $product->delete();
        $response = [
            'id' => $request->id,
            'message' => 'Producto eliminado correctamente'
        ];
        return response()->json($response, 202);
    }
}
