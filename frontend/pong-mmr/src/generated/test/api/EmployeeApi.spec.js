/*
 * tietoevry-pong-mmr
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 3.0.54
 *
 * Do not edit the class manually.
 *
 */
(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.TietoevryPongMmr);
  }
}(this, function(expect, TietoevryPongMmr) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new TietoevryPongMmr.EmployeeApi();
  });

  describe('(package)', function() {
    describe('EmployeeApi', function() {
      describe('apiEmployeesGet', function() {
        it('should call apiEmployeesGet successfully', function(done) {
          // TODO: uncomment apiEmployeesGet call and complete the assertions
          /*

          instance.apiEmployeesGet(function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            let dataCtr = data;
            expect(dataCtr).to.be.an(Array);
            expect(dataCtr).to.not.be.empty();
            for (let p in dataCtr) {
              let data = dataCtr[p];
              expect(data).to.be.a(TietoevryPongMmr.Employee);
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('apiEmployeesIdDelete', function() {
        it('should call apiEmployeesIdDelete successfully', function(done) {
          // TODO: uncomment, update parameter values for apiEmployeesIdDelete call
          /*

          instance.apiEmployeesIdDelete(id, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('apiEmployeesIdGet', function() {
        it('should call apiEmployeesIdGet successfully', function(done) {
          // TODO: uncomment, update parameter values for apiEmployeesIdGet call and complete the assertions
          /*

          instance.apiEmployeesIdGet(id, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(TietoevryPongMmr.Employee);

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('apiEmployeesIdPut', function() {
        it('should call apiEmployeesIdPut successfully', function(done) {
          // TODO: uncomment, update parameter values for apiEmployeesIdPut call
          /*

          instance.apiEmployeesIdPut(body, id, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('apiEmployeesPost', function() {
        it('should call apiEmployeesPost successfully', function(done) {
          // TODO: uncomment, update parameter values for apiEmployeesPost call and complete the assertions
          /*

          instance.apiEmployeesPost(body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(TietoevryPongMmr.Employee);

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
    });
  });

}));