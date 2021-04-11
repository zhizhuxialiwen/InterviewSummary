#ifndef VARIANTFIELD_H_
#define VARIANTFIELD_H_

#include <boost/cstdint.hpp>

namespace seamless {

union VariantField
{
    const char * strVal;
    int32_t intVal;
};

}  // namespace mlr_isearch_api

#endif // VARIANTFIELD_H_