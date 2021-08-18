<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use Illuminate\Support\Facades\DB;

class StudentController extends Controller
{

    public function createStudent(Request $request){
        $firstName = $request->input('first_name');
        $lastName = $request->input('last_name');
        $dob = $request->input('DOB');
        $data = array('first_name' => $firstName , 'last_name'=>$lastName , 'DOB' =>$dob);
        $id = DB::table('students')->insertGetId($data);

        return response()->json([
            'id' => $id,
            'first_name' => $firstName,
            'last_name' => $lastName,
            'DOB' => $dob
        ]);
    }

    public function getStudentById($id){

        $student = DB::table('students')->find($id);

        return response()->json([
            'student'=>$student


        ]);
    }
    public function getAllStudents(){

        $students = DB::table('students')->get();

        return response()->json([
            'students'=>$students

        ]);

    }
    public function deleteStudentById($id){
        DB::table('students')->delete($id);
        return response()->json([
            "message" => "student with id $id is successfully deleted!"
        ]);
    }
}
