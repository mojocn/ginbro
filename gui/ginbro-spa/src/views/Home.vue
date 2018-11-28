<template>
    <el-form ref="form" :model="form" label-width="160px">
        <el-form-item label="address">
            <el-input v-model="form.mysqlAddr" placeholder="1270.0.0.1:3306"></el-input>
        </el-form-item>
        <el-form-item label="user">
            <el-input v-model="form.mysqlUser"></el-input>
        </el-form-item>
        <el-form-item label="password">
            <el-input v-model="form.mysqlPassword"></el-input>
        </el-form-item>
        <el-form-item label="database">
            <el-input v-model="form.mysqlDatabase" placeholder="the database to create RESTful app"></el-input>
        </el-form-item>
        <el-form-item label="database">
            <el-radio v-model="form.mysqlCharset" label="utf8">utf8</el-radio>
            <el-radio v-model="form.mysqlCharset" label="utf8mb4">utf8mb4</el-radio>
        </el-form-item>
        <el-form-item label="login table">
            <el-input v-model="form.authTable" placeholder="the table for login auth"></el-input>
        </el-form-item>
        <el-form-item label="password column">
            <el-input v-model="form.authPassword" placeholder="the column for password verification"></el-input>
        </el-form-item>
        <el-form-item label="app listen">
            <el-input v-model="form.appListen" placeholder="app listening address"></el-input>
        </el-form-item>
        <el-form-item label="package">
            <el-input v-model="form.outPackage" placeholder="the path relative to $GOPATH/src"></el-input>
        </el-form-item>

        <el-form-item>
            <el-button type="primary" @click="onSubmit">Generate</el-button>
            <el-button>Cancel</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
    // @ is an alias to /src
    import HelloWorld from '@/components/HelloWorld.vue'

    export default {
        data() {
            return {
                form: {
                    mysqlAddr: '127.0.0.1:3306',
                    mysqlUser: 'root',
                    mysqlPassword: 'password',
                    mysqlDatabase: 'dbname',
                    mysqlCharset: 'utf8',
                    appListen: '127.0.0.1:5555',
                    authTable: 'users',
                    authPassword: 'password',
                    outPackage: 'ginbro-demo'
                },
                msg:""
            }
        },
        methods: {
            onSubmit() {
                mysqlGen(this.form).then(rsp => {
                    console.log(rsp)
                    this.$message({
                        message:rsp,
                        type: 'success'
                    });
                }).catch( err =>{
                    console.log(rsp)

                    this.$message({
                        message: err,
                        type: 'error'
                    });
                })
                mysqlResult().then(rsp => {
                    console.log(rsp)

                    this.$message({
                        message: rsp,
                        type: 'success'
                    });
                }).catch( err =>{

                    this.$message({
                        message: err,
                        type: 'error'
                    });
                })
            }
        }


    }
</script>
