#define APPROVALS_GOOGLETEST
#include <gtest/gtest.h>
#include "../src/TrigMath.hpp"

using namespace std;

TEST(TrigMathTest, Sin) {
    TrigMath math;
    ASSERT_NEAR(42, math.Sin(3.4), 0.000001);
}
